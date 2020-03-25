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

package json2

import (
	"testing"

	querypb "github.com/xsec-lab/go/vt/proto/query"
	vschemapb "github.com/xsec-lab/go/vt/proto/vschema"
)

func TestMarshalPB(t *testing.T) {
	col := &vschemapb.Column{
		Name: "c1",
		Type: querypb.Type_VARCHAR,
	}
	b, err := MarshalPB(col)
	if err != nil {
		t.Fatal(err)
	}
	want := "{\"name\":\"c1\",\"type\":\"VARCHAR\"}"
	got := string(b)
	if got != want {
		t.Errorf("MarshalPB(col): %q, want %q", got, want)
	}
}

func TestMarshalIndentPB(t *testing.T) {
	col := &vschemapb.Column{
		Name: "c1",
		Type: querypb.Type_VARCHAR,
	}
	b, err := MarshalIndentPB(col, "  ")
	if err != nil {
		t.Fatal(err)
	}
	want := "{\n  \"name\": \"c1\",\n  \"type\": \"VARCHAR\"\n}"
	got := string(b)
	if got != want {
		t.Errorf("MarshalPB(col): %q, want %q", got, want)
	}
}
