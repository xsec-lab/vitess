/*
Copyright 2019 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vtctld

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/net/context"

	"github.com/xsec-lab/vitess/go/vt/topo/memorytopo"
	"github.com/xsec-lab/vitess/go/vt/wrangler"

	topodatapb "github.com/xsec-lab/vitess/go/vt/proto/topodata"
	vschemapb "github.com/xsec-lab/vitess/go/vt/proto/vschema"
)

func compactJSON(in []byte) string {
	buf := &bytes.Buffer{}
	json.Compact(buf, in)
	return buf.String()
}

func TestAPI(t *testing.T) {
	ctx := context.Background()
	cells := []string{"cell1", "cell2"}
	ts := memorytopo.NewServer(cells...)
	actionRepo := NewActionRepository(ts)
	server := httptest.NewServer(nil)
	defer server.Close()

	// Populate topo. Remove ServedTypes from shards to avoid ordering issues.
	ts.CreateKeyspace(ctx, "ks1", &topodatapb.Keyspace{ShardingColumnName: "shardcol"})
	ts.CreateShard(ctx, "ks1", "-80")
	ts.CreateShard(ctx, "ks1", "80-")

	// SaveVSchema to test that creating a snapshot keyspace copies VSchema
	vs := &vschemapb.Keyspace{
		Sharded: true,
		Vindexes: map[string]*vschemapb.Vindex{
			"name1": {
				Type: "hash",
			},
		},
		Tables: map[string]*vschemapb.Table{
			"table1": {
				ColumnVindexes: []*vschemapb.ColumnVindex{
					{
						Column: "column1",
						Name:   "name1",
					},
				},
			},
		},
	}
	ts.SaveVSchema(ctx, "ks1", vs)

	tablet1 := topodatapb.Tablet{
		Alias:    &topodatapb.TabletAlias{Cell: "cell1", Uid: 100},
		Keyspace: "ks1",
		Shard:    "-80",
		Type:     topodatapb.TabletType_REPLICA,
		KeyRange: &topodatapb.KeyRange{Start: nil, End: []byte{0x80}},
		PortMap:  map[string]int32{"vt": 100},
	}
	ts.CreateTablet(ctx, &tablet1)

	tablet2 := topodatapb.Tablet{
		Alias:    &topodatapb.TabletAlias{Cell: "cell2", Uid: 200},
		Keyspace: "ks1",
		Shard:    "-80",
		Type:     topodatapb.TabletType_REPLICA,
		KeyRange: &topodatapb.KeyRange{Start: nil, End: []byte{0x80}},
		PortMap:  map[string]int32{"vt": 200},
	}
	ts.CreateTablet(ctx, &tablet2)

	// Populate fake actions.
	actionRepo.RegisterKeyspaceAction("TestKeyspaceAction",
		func(ctx context.Context, wr *wrangler.Wrangler, keyspace string, r *http.Request) (string, error) {
			return "TestKeyspaceAction Result", nil
		})
	actionRepo.RegisterShardAction("TestShardAction",
		func(ctx context.Context, wr *wrangler.Wrangler, keyspace, shard string, r *http.Request) (string, error) {
			return "TestShardAction Result", nil
		})
	actionRepo.RegisterTabletAction("TestTabletAction", "",
		func(ctx context.Context, wr *wrangler.Wrangler, tabletAlias *topodatapb.TabletAlias, r *http.Request) (string, error) {
			return "TestTabletAction Result", nil
		})

	realtimeStats := newRealtimeStatsForTesting()
	initAPI(ctx, ts, actionRepo, realtimeStats)

	ts1 := tabletStats("ks1", "cell1", "-80", topodatapb.TabletType_REPLICA, 100)
	ts2 := tabletStats("ks1", "cell1", "-80", topodatapb.TabletType_RDONLY, 200)
	ts3 := tabletStats("ks1", "cell2", "80-", topodatapb.TabletType_REPLICA, 300)
	ts4 := tabletStats("ks1", "cell2", "80-", topodatapb.TabletType_RDONLY, 400)

	ts5 := tabletStats("ks2", "cell1", "0", topodatapb.TabletType_REPLICA, 500)
	ts6 := tabletStats("ks2", "cell2", "0", topodatapb.TabletType_REPLICA, 600)

	realtimeStats.StatsUpdate(ts1)
	realtimeStats.StatsUpdate(ts2)
	realtimeStats.StatsUpdate(ts3)
	realtimeStats.StatsUpdate(ts4)
	realtimeStats.StatsUpdate(ts5)
	realtimeStats.StatsUpdate(ts6)

	// Test cases.
	table := []struct {
		method, path, body, want string
	}{
		// Create snapshot keyspace using API
		{"POST", "vtctl/", `["CreateKeyspace", "-keyspace_type=SNAPSHOT", "-base_keyspace=ks1", "-snapshot_time=2006-01-02T15:04:05+00:00", "ks3"]`, `{
		   "Error": "",
		   "Output": ""
		}`},

		// Cells
		{"GET", "cells", "", `["cell1","cell2"]`},

		// Keyspaces
		{"GET", "keyspaces", "", `["ks1", "ks3"]`},
		{"GET", "keyspaces/ks1", "", `{
				"sharding_column_name": "shardcol",
				"sharding_column_type": 0,
				"served_froms": [],
                                "keyspace_type":0,
                                "base_keyspace":"",
                                "snapshot_time":null
			}`},
		{"GET", "keyspaces/nonexistent", "", "404 page not found"},
		{"POST", "keyspaces/ks1?action=TestKeyspaceAction", "", `{
				"Name": "TestKeyspaceAction",
				"Parameters": "ks1",
				"Output": "TestKeyspaceAction Result",
				"Error": false
			}`},

		// Shards
		{"GET", "shards/ks1/", "", `["-80","80-"]`},
		{"GET", "shards/ks1/-80", "", `{
				"master_alias": null,
				"master_term_start_time":null,
				"key_range": {
					"start": null,
					"end":"gA=="
				},
				"served_types": [],
				"source_shards": [],
				"tablet_controls": [],
				"is_master_serving": true
			}`},
		{"GET", "shards/ks1/-DEAD", "", "404 page not found"},
		{"POST", "shards/ks1/-80?action=TestShardAction", "", `{
				"Name": "TestShardAction",
				"Parameters": "ks1/-80",
				"Output": "TestShardAction Result",
				"Error": false
			}`},

		// Tablets
		{"GET", "tablets/?shard=ks1%2F-80", "", `[
				{"cell":"cell1","uid":100},
				{"cell":"cell2","uid":200}
			]`},
		{"GET", "tablets/?cell=cell1", "", `[
				{"cell":"cell1","uid":100}
			]`},
		{"GET", "tablets/?shard=ks1%2F-80&cell=cell2", "", `[
				{"cell":"cell2","uid":200}
			]`},
		{"GET", "tablets/?shard=ks1%2F80-&cell=cell1", "", `[]`},
		{"GET", "tablets/cell1-100", "", `{
				"alias": {"cell": "cell1", "uid": 100},
				"hostname": "",
				"port_map": {"vt": 100},
				"keyspace": "ks1",
				"shard": "-80",
				"key_range": {
					"start": null,
					"end": "gA=="
				},
				"type": 2,
				"db_name_override": "",
				"tags": {},
				"mysql_hostname":"",
				"mysql_port":0,
				"master_term_start_time":null
			}`},
		{"GET", "tablets/nonexistent-999", "", "404 page not found"},
		{"POST", "tablets/cell1-100?action=TestTabletAction", "", `{
				"Name": "TestTabletAction",
				"Parameters": "cell1-0000000100",
				"Output": "TestTabletAction Result",
				"Error": false
			}`},

		// Tablet Updates
		{"GET", "tablet_statuses/?keyspace=ks1&cell=cell1&type=REPLICA&metric=lag", "", `[
		{
		    "Data": [ [100, -1] ],
		    "Aliases": [[ { "cell": "cell1", "uid": 100 }, null ]],
		    "KeyspaceLabel": { "Name": "ks1", "Rowspan": 1 },
		    "CellAndTypeLabels": [{ "CellLabel": { "Name": "cell1",  "Rowspan": 1 }, "TypeLabels": [{"Name": "REPLICA", "Rowspan": 1}] }] ,
		    "ShardLabels": ["-80", "80-"],
		    "YGridLines": [0.5]
		  }
		]`},
		{"GET", "tablet_statuses/?keyspace=ks1&cell=all&type=all&metric=lag", "", `[
		{
		  "Data":[[-1,400],[-1,300],[200,-1],[100,-1]],
		  "Aliases":[[null,{"cell":"cell2","uid":400}],[null,{"cell":"cell2","uid":300}],[{"cell":"cell1","uid":200},null],[{"cell":"cell1","uid":100},null]],
		  "KeyspaceLabel":{"Name":"ks1","Rowspan":4},
		  "CellAndTypeLabels":[
		     {"CellLabel":{"Name":"cell1","Rowspan":2},"TypeLabels":[{"Name":"REPLICA","Rowspan":1},{"Name":"RDONLY","Rowspan":1}]},
		     {"CellLabel":{"Name":"cell2","Rowspan":2},"TypeLabels":[{"Name":"REPLICA","Rowspan":1},{"Name":"RDONLY","Rowspan":1}]}],
		  "ShardLabels":["-80","80-"],
		  "YGridLines":[0.5,1.5,2.5,3.5]
		}
		]`},
		{"GET", "tablet_statuses/?keyspace=all&cell=all&type=all&metric=lag", "", `[
		  {
		   "Data":[[-1,300],[200,-1]],
		   "Aliases":null,
		   "KeyspaceLabel":{"Name":"ks1","Rowspan":2},
		  "CellAndTypeLabels":[
		    {"CellLabel":{"Name":"cell1","Rowspan":1},"TypeLabels":null},
		    {"CellLabel":{"Name":"cell2","Rowspan":1},"TypeLabels":null}],
		  "ShardLabels":["-80","80-"],
		  "YGridLines":[0.5,1.5]
		  },
		  {
		    "Data":[[600],[500]],
		   "Aliases":null,
		   "KeyspaceLabel":{"Name":"ks2","Rowspan":2},
		  "CellAndTypeLabels":[
		    {"CellLabel":{"Name":"cell1","Rowspan":1},"TypeLabels":null},
		    {"CellLabel":{"Name":"cell2","Rowspan":1},"TypeLabels":null}],
		  "ShardLabels":["0"],
		  "YGridLines":[0.5, 1.5]
		  }
		]`},
		{"GET", "tablet_statuses/cell1/REPLICA/lag", "", "can't get tablet_statuses: invalid target path: \"cell1/REPLICA/lag\"  expected path: ?keyspace=<keyspace>&cell=<cell>&type=<type>&metric=<metric>"},
		{"GET", "tablet_statuses/?keyspace=ks1&cell=cell1&type=hello&metric=lag", "", "can't get tablet_statuses: invalid tablet type: unknown TabletType hello"},

		// Tablet Health
		{"GET", "tablet_health/cell1/100", "", `{ "Key": "", "Tablet": { "alias": { "cell": "cell1", "uid": 100 },"port_map": { "vt": 100 }, "keyspace": "ks1", "shard": "-80", "type": 2},
		  "Name": "", "Target": { "keyspace": "ks1", "shard": "-80", "tablet_type": 2 }, "Up": true, "Serving": true, "TabletExternallyReparentedTimestamp": 0,
		  "Stats": { "seconds_behind_master": 100 }, "LastError": null }`},
		{"GET", "tablet_health/cell1", "", "can't get tablet_health: invalid tablet_health path: \"cell1\"  expected path: /tablet_health/<cell>/<uid>"},
		{"GET", "tablet_health/cell1/gh", "", "can't get tablet_health: incorrect uid"},

		// Topology Info
		{"GET", "topology_info/?keyspace=all&cell=all", "", `{
		   "Keyspaces": ["ks1", "ks2"],
		   "Cells": ["cell1","cell2"],
		   "TabletTypes": ["REPLICA","RDONLY"]
		}`},
		{"GET", "topology_info/?keyspace=ks1&cell=cell1", "", `{
		   "Keyspaces": ["ks1", "ks2"],
		   "Cells": ["cell1","cell2"],
		   "TabletTypes": ["REPLICA", "RDONLY"]
		}`},

		// vtctl RunCommand
		{"POST", "vtctl/", `["GetKeyspace","ks1"]`, `{
		   "Error": "",
		   "Output": "{\n  \"sharding_column_name\": \"shardcol\",\n  \"sharding_column_type\": 0,\n  \"served_froms\": [\n  ],\n  \"keyspace_type\": 0,\n  \"base_keyspace\": \"\",\n  \"snapshot_time\": null\n}\n\n"
		}`},
		{"POST", "vtctl/", `["GetKeyspace","ks3"]`, `{
		   "Error": "",
		   "Output": "{\n  \"sharding_column_name\": \"\",\n  \"sharding_column_type\": 0,\n  \"served_froms\": [\n  ],\n  \"keyspace_type\": 1,\n  \"base_keyspace\": \"ks1\",\n  \"snapshot_time\": {\n    \"seconds\": \"1136214245\",\n    \"nanoseconds\": 0\n  }\n}\n\n"
		}`},
		{"POST", "vtctl/", `["GetVSchema","ks3"]`, `{
		   "Error": "",
		   "Output": "{\n  \"sharded\": true,\n  \"vindexes\": {\n    \"name1\": {\n      \"type\": \"hash\"\n    }\n  },\n  \"tables\": {\n    \"table1\": {\n      \"columnVindexes\": [\n        {\n          \"column\": \"column1\",\n          \"name\": \"name1\"\n        }\n      ]\n    }\n  },\n  \"requireExplicitRouting\": true\n}\n\n"
		}`},
		{"POST", "vtctl/", `["GetKeyspace","does_not_exist"]`, `{
			"Error": "node doesn't exist: keyspaces/does_not_exist/Keyspace",
		   "Output": ""
		}`},
		{"POST", "vtctl/", `["Panic"]`, `uncaught panic: this command panics on purpose`},
	}
	for _, in := range table {
		t.Run(in.method+in.path, func(t *testing.T) {
			var resp *http.Response
			var err error

			switch in.method {
			case "GET":
				resp, err = http.Get(server.URL + apiPrefix + in.path)
			case "POST":
				resp, err = http.Post(server.URL+apiPrefix+in.path, "application/json", strings.NewReader(in.body))
			default:
				t.Fatalf("[%v] unknown method: %v", in.path, in.method)
				return
			}

			if err != nil {
				t.Fatalf("[%v] http error: %v", in.path, err)
				return
			}

			body, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()

			if err != nil {
				t.Fatalf("[%v] ioutil.ReadAll(resp.Body) error: %v", in.path, err)
				return
			}

			got := compactJSON(body)
			want := compactJSON([]byte(in.want))
			if want == "" {
				// want is not valid JSON. Fallback to a string comparison.
				want = in.want
				// For unknown reasons errors have a trailing "\n\t\t". Remove it.
				got = strings.TrimSpace(string(body))
			}
			if !strings.HasPrefix(got, want) {
				t.Fatalf("For path [%v] got\n'%v', want\n'%v'", in.path, got, want)
				return
			}
		})

	}
}
