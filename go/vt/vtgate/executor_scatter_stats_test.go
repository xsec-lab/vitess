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

package vtgate

import (
	"context"
	"testing"

	"net/http/httptest"

	"github.com/stretchr/testify/require"

	vtgatepb "github.com/xsec-lab/go/vt/proto/vtgate"
)

func TestScatterStatsWithNoScatterQuery(t *testing.T) {
	executor, _, _, _ := createExecutorEnv()
	session := NewSafeSession(&vtgatepb.Session{TargetString: "@master"})

	_, err := executor.Execute(context.Background(), "TestExecutorResultsExceeded", session, "select * from main1", nil)
	require.NoError(t, err)

	result, err := executor.gatherScatterStats()
	require.NoError(t, err)
	require.Equal(t, 0, len(result.Items))
}

func TestScatterStatsWithSingleScatterQuery(t *testing.T) {
	executor, _, _, _ := createExecutorEnv()
	session := NewSafeSession(&vtgatepb.Session{TargetString: "@master"})

	_, err := executor.Execute(context.Background(), "TestExecutorResultsExceeded", session, "select * from user", nil)
	require.NoError(t, err)

	result, err := executor.gatherScatterStats()
	require.NoError(t, err)
	require.Equal(t, 1, len(result.Items))
}

func TestScatterStatsHttpWriting(t *testing.T) {
	executor, _, _, _ := createExecutorEnv()
	session := NewSafeSession(&vtgatepb.Session{TargetString: "@master"})

	_, err := executor.Execute(context.Background(), "TestExecutorResultsExceeded", session, "select * from user", nil)
	require.NoError(t, err)

	_, err = executor.Execute(context.Background(), "TestExecutorResultsExceeded", session, "select * from user where Id = 15", nil)
	require.NoError(t, err)

	_, err = executor.Execute(context.Background(), "TestExecutorResultsExceeded", session, "select * from user where Id > 15", nil)
	require.NoError(t, err)

	query4 := "select * from user as u1 join  user as u2 on u1.Id = u2.Id"
	_, err = executor.Execute(context.Background(), "TestExecutorResultsExceeded", session, query4, nil)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()
	executor.WriteScatterStats(recorder)

	// Here we are checking that the template was executed correctly.
	// If it wasn't, instead of html, we'll get an error message
	require.Contains(t, recorder.Body.String(), query4)
	require.NoError(t, err)
}
