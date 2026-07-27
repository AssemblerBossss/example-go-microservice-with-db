package sqlite_clock

import (
	"database/sql"
	"time"
)

type SQLiteClock struct {
	DB *sql.DB
}

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

func (c *SQLiteClock) getUnixSeconds() int64 {
	var unixSeconds int64
	if err := c.DB.QueryRow(`SELECT CAST(strftime('%s','now' ) AS INTEGER);`).Scan(&unixSeconds); err != nil {
		panic(err)
	}
	return unixSeconds
}

func (c *SQLiteClock) NowUnix() int64 {
	//t := c.getTimeFromCurrentTimestamp()
	//return t.UTC().Unix()
	return c.getUnixSeconds()
}

func NewSQLiteClock(db *sql.DB) *SQLiteClock {
	return &SQLiteClock{db}
}
