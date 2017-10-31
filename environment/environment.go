/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package environment

// Environment ...
type Environment struct {
	ID          int                    `json:"id"`
	Name        string                 `json:"name"`
	Type        string                 `json:"type"`
	Status      string                 `json:"status"`
	Options     map[string]interface{} `json:"options"`
	Credentials map[string]interface{} `json:"credentials"`
}

// Ready : returns true if the environment is not busy
func (e *Environment) Ready() bool {
	switch e.Status {
	case "done":
		return true
	case "errored", "initializing", "in_progress":
		return false
	default:
		return false
	}
}

// SyncEnabled : returns true if sync is enabled for an environment
func (e *Environment) SyncEnabled() bool {
	if e.Options["sync"] != nil {
		return e.Options["sync"].(bool)
	}
	return false
}
