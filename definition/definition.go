/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

// Definition ...
type Definition map[string]interface{}

// Env : definition env
func (d *Definition) Name() string {
	name, _ := (*d)["name"].(string)
	return name
}

// Project : definition project
func (d *Definition) Project() string {
	project, _ := (*d)["project"].(string)
	return project
}

// FullName : full env name (project + env)
func (d *Definition) FullName() string {
	return d.Project() + "/" + d.Name()
}
