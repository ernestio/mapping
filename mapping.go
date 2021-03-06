/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapping

import (
	"errors"

	"github.com/ernestio/mapping/build"
	"github.com/ernestio/mapping/definition"
	"github.com/ernestio/mapping/environment"
	"github.com/ernestio/mapping/policy"
	"github.com/ernestio/mapping/query"
	"github.com/ernestio/mapping/validation"
	"github.com/r3labs/akira"
	"github.com/satori/uuid"
)

// Mapping : stores a environments build mapping
type Mapping struct {
	Environment string
	Changelog   bool
	Result      map[string]interface{}
	Validation  validation.Validation
	conn        akira.Connector
}

// New : create a new mapping
func New(c akira.Connector, env string) *Mapping {
	return &Mapping{
		Environment: env,
		conn:        c,
	}
}

// Validate : validate a build mapping against an environments policy documents
func (m *Mapping) Validate() error {
	var policies []policy.Policy
	var documents []policy.PolicyDocument

	q := map[string][]string{"environments": []string{m.Environment}}

	err := query.New(m.conn, "policy.find").Request(q).Run(&policies)
	if err != nil {
		return err
	}

	if len(policies) < 1 {
		return nil
	}

	for _, p := range policies {
		var pd []policy.PolicyDocument

		pq := map[string]int{"policy_id": p.ID}

		err := query.New(m.conn, "policy_document.find").Request(pq).Run(&pd)
		if err != nil {
			return err
		}

		if len(pd) < 1 {
			continue
		}

		documents = append(documents, pd[0])
	}

	bv := validation.BuildValidation{
		Mapping:  m.Result,
		Policies: documents,
	}

	return query.New(m.conn, "build.validate").Request(bv).Run(&m.Validation)
}

// Diff : gets a mapping for a diff between two environment builds
func (m *Mapping) Diff(a, b string) error {
	var err error
	var ag, bg map[string]interface{}

	err = query.New(m.conn, "build.get.mapping").ID(a).Run(&ag)
	if err != nil {
		return err
	}

	err = query.New(m.conn, "build.get.mapping").ID(b).Run(&bg)
	if err != nil {
		return err
	}

	return m.DiffGraphs(ag, bg)
}

// DiffGraphs : gets a mpping for a diff between two graphs
func (m *Mapping) DiffGraphs(ag, bg map[string]interface{}) error {
	var err error
	credentials, err := GetCredentials(m.conn, m.Environment)
	if err != nil {
		return err
	}
	r := Request{
		ID:          uuid.NewV4().String(),
		Name:        m.Environment,
		From:        ag,
		To:          bg,
		Credentials: credentials,
	}

	return query.New(m.conn, "mapping.get.diff").Request(&r).Run(&m.Result)
}

// Import : gets a mapping for import
func (m *Mapping) Import(filters []string) error {
	credentials, err := GetCredentials(m.conn, m.Environment)
	if err != nil {
		return err
	}

	r := Request{
		ID:          uuid.NewV4().String(),
		Name:        m.Environment,
		Filters:     filters,
		Credentials: credentials,
	}

	return query.New(m.conn, "mapping.get.import").Request(&r).Run(&m.Result)
}

// Apply : apply a definition
func (m *Mapping) Apply(d *definition.Definition) error {
	var env environment.Environment
	var builds []build.Build

	err := query.New(m.conn, "environment.get").Name(m.Environment).Run(&env)
	if err != nil {
		return err
	}

	err = query.New(m.conn, "build.find").Filter("environment_id", env.ID).Run(&builds)
	if err != nil {
		return err
	}

	if len(builds) > 0 {
		return m.update(d, builds[0].UUID)
	}

	return m.create(d)
}

// Delete : gets a mapping for deleting an environment
func (m *Mapping) Delete() error {
	var env environment.Environment
	var builds []build.Build
	var mapping map[string]interface{}

	err := query.New(m.conn, "environment.get").Name(m.Environment).Run(&env)
	if err != nil {
		return err
	}

	err = query.New(m.conn, "build.find").Filter("environment_id", env.ID).Run(&builds)
	if err != nil {
		return err
	}

	if len(builds) < 1 {
		return errors.New("environment has no builds")
	}

	err = query.New(m.conn, "build.get.mapping").ID(builds[0].UUID).Run(&mapping)
	if err != nil {
		return err
	}

	credentials, err := GetCredentials(m.conn, m.Environment)
	if err != nil {
		return err
	}

	r := Request{
		ID:          uuid.NewV4().String(),
		Name:        m.Environment,
		From:        mapping,
		Credentials: credentials,
	}

	return query.New(m.conn, "mapping.get.delete").Request(&r).Run(&m.Result)
}

// Create : gets a mapping for creating an environment
func (m *Mapping) create(d *definition.Definition) error {
	credentials, err := GetCredentials(m.conn, m.Environment)
	if err != nil {
		return err
	}

	r := Request{
		ID:          uuid.NewV4().String(),
		Name:        m.Environment,
		Definition:  *d,
		Credentials: credentials,
		Changelog:   m.Changelog,
	}

	return query.New(m.conn, "mapping.get.create").Request(&r).Run(&m.Result)
}

// Update : gets a mapping for updating an existing environment
func (m *Mapping) update(d *definition.Definition, build string) error {
	var mapping map[string]interface{}

	credentials, err := GetCredentials(m.conn, m.Environment)
	if err != nil {
		return err
	}

	err = query.New(m.conn, "build.get.mapping").ID(build).Run(&mapping)
	if err != nil {
		return err
	}

	r := Request{
		ID:          uuid.NewV4().String(),
		Name:        m.Environment,
		From:        mapping,
		Definition:  *d,
		Credentials: credentials,
		Changelog:   m.Changelog,
	}

	return query.New(m.conn, "mapping.get.update").Request(&r).Run(&m.Result)
}
