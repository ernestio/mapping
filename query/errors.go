/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package query

import (
	"encoding/json"
	"errors"
)

// ErrNotFound : could not find a resource
var ErrNotFound = errors.New("not found")

// Error : Error json structure
type Error struct {
	Message string `json:"_error"`
}

func validate(data []byte) error {
	var e Error
	_ = json.Unmarshal(data, &e)
	if e.Message != "" {
		return errors.New(e.Message)
	}
	return nil
}
