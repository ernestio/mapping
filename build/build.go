/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package build

// Build ...
type Build struct {
	ID            uint                   `json:"-" gorm:"primary_key"`
	UUID          string                 `json:"id"`
	EnvironmentID uint                   `json:"environment_id"`
	UserID        uint                   `json:"user_id"`
	Type          string                 `json:"type"`
	Status        string                 `json:"status"`
	Definition    string                 `json:"definition" gorm:"type:text;"`
	Mapping       map[string]interface{} `json:"mapping"`
}
