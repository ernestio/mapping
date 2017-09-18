/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package query

import (
	"errors"
	"strings"
)

// ErrNotFound : could not find a resource
var ErrNotFound = errors.New("not found")

func isError(data []byte, err error) bool {
	return strings.Contains(string(data), err.Error())
}
