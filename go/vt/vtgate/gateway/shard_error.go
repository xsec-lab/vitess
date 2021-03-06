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

package gateway

import (
	"github.com/xsec-lab/go/vt/topo/topoproto"
	"github.com/xsec-lab/go/vt/topotools"
	"github.com/xsec-lab/go/vt/vterrors"

	querypb "github.com/xsec-lab/go/vt/proto/query"
	topodatapb "github.com/xsec-lab/go/vt/proto/topodata"
)

// NewShardError returns a new error with the shard info amended.
func NewShardError(in error, target *querypb.Target, tablet *topodatapb.Tablet) error {
	if in == nil {
		return nil
	}
	if tablet != nil {
		return vterrors.Wrapf(in, "target: %s.%s.%s, used tablet: %s", target.Keyspace, target.Shard, topoproto.TabletTypeLString(target.TabletType), topotools.TabletIdent(tablet))
	}
	if target != nil {
		return vterrors.Wrapf(in, "target: %s.%s.%s", target.Keyspace, target.Shard, topoproto.TabletTypeLString(target.TabletType))
	}
	return in
}
