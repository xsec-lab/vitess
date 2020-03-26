package main

// This plugin imports opentsdb to register the opentsdb stats backend.

import (
	"github.com/xsec-lab/vitess/go/stats/opentsdb"
)

func init() {
	opentsdb.Init("vtctld")
}
