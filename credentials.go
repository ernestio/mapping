/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapping

import (
	"strings"

	"github.com/ernestio/mapping/environment"
	"github.com/ernestio/mapping/project"
	"github.com/ernestio/mapping/query"
	"github.com/r3labs/akira"
)

// GetCredentials : gets credentials from a project
func GetCredentials(c akira.Connector, env string) (map[string]interface{}, error) {
	var p project.Project
	var e environment.Environment

	err := query.New(c, "datacenter.get").Name(projectname(env)).Run(&p)
	if err != nil {
		return nil, err
	}

	err = query.New(c, "environment.get").Name(env).Run(&e)
	if err != nil {
		return nil, err
	}

	p.Override(e.Credentials)

	return p.Credentials, nil
}

func projectname(env string) string {
	return strings.Split(env, "/")[0]
}
