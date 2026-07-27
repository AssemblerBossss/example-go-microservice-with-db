package sqlite_clock

import (
	"database/sql"
	"time"
)

type SQLiteClock struct {
	DB *sql.DB
}

// getTimeFromCurrentTimestamp returns current time from SQLite as time.Time.
func (c *SQLiteClock) getTimeFromCurrentTimestamp() time.Time {
	var ts string
	if err := c.DB.QueryRow(`SELECT CURRENT_TIMESTAMP;`).Scan(&ts); err != nil {
		panic(err)
	}
	t, err := time.Parse(time.DateTime, ts)
	if err != nil {
		panic(err)
	}
	return t
}

// getUnixSeconds returns current Unix timestamp directly from SQLite using strftime.
func (c *SQLiteClock) getUnixSeconds() int64 {
	var unixSeconds int64
	if err := c.DB.QueryRow(`SELECT CAST(strftime('%s','now' ) AS INTEGER);`).Scan(&unixSeconds); err != nil {
		panic(err)
	}
	return unixSeconds
}

// NowUnix returns current Unix timestamp from SQLite. Implements domains.Clock.
func (c *SQLiteClock) NowUnix() int64 {
	//t := c.getTimeFromCurrentTimestamp()
	//return t.UTC().Unix()
	return c.getUnixSeconds()
}

// NewSQLiteClock creates a new SQLiteClock with the given database connection.
func NewSQLiteClock(db *sql.DB) *SQLiteClock {
	return &SQLiteClock{db}
}
