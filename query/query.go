/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package query

import (
	"encoding/json"
	"time"

	"github.com/r3labs/akira"
)

// Query ...
type Query struct {
	Connection akira.Connector
	subject    string
	result     interface{}
	request    interface{}
	params     map[string]interface{}
}

// New : creates a new query
func New(conn akira.Connector, subject string) *Query {
	return &Query{
		Connection: conn,
		subject:    subject,
	}
}

// Run : runs the query
func (q *Query) Run(model interface{}) error {
	q.result = model
	return q.query()
}

func (q *Query) build() ([]byte, error) {
	if q.params != nil {
		return json.Marshal(q.params)
	}

	if q.request != nil {
		return json.Marshal(q.request)
	}

	return json.Marshal(q.result)
}

func (q *Query) query() error {
	data, err := q.build()
	if err != nil {
		return err
	}

	resp, err := q.Connection.Request(q.subject, data, time.Second)
	if err != nil {
		return err
	}

	if q.result == nil {
		return nil
	}

	err = validate(resp.Data)
	if err != nil {
		return err
	}

	return json.Unmarshal(resp.Data, q.result)
}
