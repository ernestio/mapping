/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapping

import (
	"errors"

	"github.com/ernestio/mapping/build"
	"github.com/ernestio/mapping/definition"
	"github.com/ernestio/mapping/environment"
	"github.com/ernestio/mapping/query"
	"github.com/r3labs/akira"
	"github.com/satori/uuid"
)

// Mapping : stores a environments build mapping
type Mapping struct {
	Environment string
	Result      map[string]interface{}
	conn        akira.Connector
}

// New : create a new mapping
func New(c akira.Connector, env string) *Mapping {
	return &Mapping{
		Environment: env,
		conn:        c,
	}
}

// Diff : gets a mapping for a diff between two environment builds
func (m *Mapping) Diff(a, b string) error {
	return nil
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

	err = query.New(m.conn, "builds.find").Filter("environment_id", env.ID).Run(&builds)
	if err != nil {
		return err
	}

	if len(builds) > 1 {
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

	err = query.New(m.conn, "builds.find").Filter("environment_id", env.ID).Run(&builds)
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
	}

	return query.New(m.conn, "mapping.get.update").Request(&r).Run(&m.Result)
}
