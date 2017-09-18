/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

import "strings"

// Definition ...
type Definition map[string]interface{}

// Env : returns the environment name
func (d *Definition) Env() string {
	name, _ := (*d)["name"].(string)
	return name
}

// Project : returns the project name
func (d *Definition) Project() string {
	return strings.Split(d.Env(), "/")[0]
}
