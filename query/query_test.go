/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package query

import (
	"testing"

	"github.com/ernestio/mapping/environment"
	"github.com/r3labs/akira"
	"github.com/stretchr/testify/assert"
)

var fc akira.Connector

func TestQuery(t *testing.T) {
	cases := []struct {
		Name, Subject string
		Filter        map[string]interface{}
		Model         interface{}
		Expected      *environment.Environment
	}{
		{"get", "environment.get", map[string]interface{}{"id": 1}, &environment.Environment{}, &environment.Environment{ID: 1, Name: "env-1"}},
		{"find", "environment.find", map[string]interface{}{"name": "env-2"}, &[]environment.Environment{}, &environment.Environment{ID: 2, Name: "env-2"}},
		{"create", "environment.set", nil, &environment.Environment{Name: "env-3"}, &environment.Environment{ID: 3, Name: "env-3"}},
		{"update", "environment.set", nil, &environment.Environment{ID: 2, Name: "env-2", Status: "something"}, &environment.Environment{ID: 2, Name: "env-2", Status: "something"}},
		{"delete", "environment.del", map[string]interface{}{"id": 1}, nil, nil},
	}

	fc = akira.NewFakeConnector()
	xfc := fc.(*akira.FakeConnector)

	fc.Subscribe("environment.get", testGetHandler)
	fc.Subscribe("environment.set", testSetHandler)
	fc.Subscribe("environment.del", testDeleteHandler)
	fc.Subscribe("environment.find", testFindHandler)

	for _, tc := range cases {
		xfc.ResetEvents()

		t.Run(tc.Name, func(t *testing.T) {
			q := New(fc, tc.Subject)
			for k, v := range tc.Filter {
				q = q.Filter(k, v)
			}

			err := q.Run(tc.Model)

			assert.Nil(t, err)
			assert.Equal(t, 1, len(xfc.Events[tc.Subject]))

			switch qr := tc.Model.(type) {
			case *[]environment.Environment:
				assert.Equal(t, 1, len(*qr))
				assert.Equal(t, tc.Expected.ID, (*qr)[0].ID)
				assert.Equal(t, tc.Expected.Name, (*qr)[0].Name)
			case *environment.Environment:
				assert.Equal(t, tc.Expected.ID, qr.ID)
				assert.Equal(t, tc.Expected.Name, qr.Name)
			}
		})
	}

}
