/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapping

// Request :
type Request struct {
	ID          string                 `json:"id,omitempty"`
	Name        string                 `json:"name,omitempty"`
	UserID      int                    `json:"user_id"`
	Username    string                 `json:"username"`
	Filters     []string               `json:"filters,omitempty"`
	Definition  map[string]interface{} `json:"definition,omitempty"`
	From        map[string]interface{} `json:"from,omitempty"`
	To          map[string]interface{} `json:"to,omitempty"`
	Credentials map[string]interface{} `json:"credentials,omitempty"`
}
