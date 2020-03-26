/*
Copyright 2017 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreedto in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package events

import topodatapb "github.com/xsec-lab/go/vt/proto/topodata"

// TabletChange is an event that describes changes to a tablet's topo record.
// It is triggered when the CURRENT process changes ANY tablet's record.
// It is NOT triggered when a DIFFERENT process changes THIS tablet's record.
// To be notified when THIS tablet's record changes, even if it was changed
// by a different process, listen for go/vt/tabletmanager/events.StateChange.
type TabletChange struct {
	Tablet topodatapb.Tablet
	Status string
}
