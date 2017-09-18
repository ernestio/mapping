/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package project

// Project ...
type Project struct {
	ID          uint                   `json:"id"`
	Name        string                 `json:"name"`
	Type        string                 `json:"type"`
	Credentials map[string]interface{} `json:"credentials"`
}

// Override : override project credentials
func (p *Project) Override(credentials map[string]interface{}) {
	for k, v := range credentials {
		p.Credentials[k] = v
	}
}
