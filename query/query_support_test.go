/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package query

import (
	"encoding/json"

	"github.com/ernestio/mapping/environment"
	"github.com/nats-io/nats"
)

var testData = []environment.Environment{
	{ID: 1, Name: "env-1"},
	{ID: 2, Name: "env-2"},
}

func testFindHandler(m *nats.Msg) {
	q := map[string]interface{}{}
	json.Unmarshal(m.Data, &q)

	results := []environment.Environment{}

	for _, td := range testData {
		if q["name"] != nil {
			if td.Name == q["name"].(string) {
				results = append(results, td)
			}
		}
	}

	data, _ := json.Marshal(results)
	fc.Publish(m.Reply, data)
}

func testGetHandler(m *nats.Msg) {
	q := map[string]interface{}{}
	json.Unmarshal(m.Data, &q)

	for _, td := range testData {
		if q["id"] != nil {
			if td.ID == int(q["id"].(float64)) {
				data, _ := json.Marshal(td)
				fc.Publish(m.Reply, data)
				return
			}
		}
	}
}

func testSetHandler(m *nats.Msg) {
	e := environment.Environment{}
	json.Unmarshal(m.Data, &e)

	if e.ID == 0 {
		e.ID = int(len(testData)) + 1
		data, _ := json.Marshal(e)
		fc.Publish(m.Reply, data)
		return
	}

	for i := 0; i < len(testData); i++ {
		if e.ID == testData[i].ID {
			testData[i] = e
			data, _ := json.Marshal(e)
			fc.Publish(m.Reply, data)
			return
		}
	}
}

func testDeleteHandler(m *nats.Msg) {
	fc.Publish(m.Reply, []byte(`{"status": "success"}`))
}
