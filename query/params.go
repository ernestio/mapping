/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package query

// Name : filter by name
func (q *Query) Name(name string) *Query {
	q.setParam("name", name)
	return q
}

// ID : filter by id
func (q *Query) ID(id string) *Query {
	q.setParam("id", id)
	return q
}

// Filter : filter by a custom key/value pair
func (q *Query) Filter(k string, v interface{}) *Query {
	q.setParam(k, v)
	return q
}

// Request : sets the request object
func (q *Query) Request(r interface{}) *Query {
	q.request = r
	return q
}

func (q *Query) setParam(k string, v interface{}) {
	if q.params == nil {
		q.params = make(map[string]interface{})
	}
	q.params[k] = v
}
