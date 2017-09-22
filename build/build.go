/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package build

// Build ...
type Build struct {
	ID            int                    `json:"-" gorm:"primary_key"`
	UUID          string                 `json:"id"`
	EnvironmentID int                    `json:"environment_id"`
	UserID        int                    `json:"user_id"`
	Username      string                 `json:"user_name"`
	Type          string                 `json:"type"`
	Status        string                 `json:"status"`
	Definition    string                 `json:"definition" gorm:"type:text;"`
	Mapping       map[string]interface{} `json:"mapping"`
}
