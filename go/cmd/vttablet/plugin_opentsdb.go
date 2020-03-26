package main

// This plugin imports opentsdb to register the opentsdb stats backend.

import (
	"github.com/xsec-lab/go/stats/opentsdb"
)

func init() {
	opentsdb.Init("vttablet")
}
