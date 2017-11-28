// Package reap contains the history reaping subsystem for horizon.  This system
// is designed to remove data from the history database such that it does not
// grow indefinitely.  The system can be configured with a number of ledgers to
// maintain at a minimum.
package reap

import (
	"gitlab.com/swarmfund/horizon/db2"
	"time"
)

// System represents the history reaping subsystem of horizon.
type System struct {
	HorizonDB      *db2.Repo
	RetentionCount uint

	nextRun time.Time
}

// New initializes the reaper, causing it to begin polling the stellar-core
// database for now ledgers and ingesting data into the horizon database.
func New(retention uint, horizon *db2.Repo) *System {
	r := &System{
		HorizonDB:      horizon,
		RetentionCount: retention,
	}

	r.nextRun = time.Now().Add(1 * time.Hour)
	return r
}
