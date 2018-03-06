/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package policy

// Policy : stores policy data
type Policy struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Definition   string   `json:"definition"`
	Environments []string `json:"environments"`
}
