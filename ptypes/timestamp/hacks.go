package timestamp

import (
	"errors"
	"time"

	"database/sql/driver"

	"fmt"
)

// Conform to the Scanner interface for database/sql
func (m *Timestamp) Scan(value interface{}) error {
	t, ok := value.(time.Time)
	if !ok {
		return errors.New("incompatible type passed, expected time.Time")
	}

	seconds := t.Unix()
	nanos := int32(t.Sub(time.Unix(seconds, 0)))
	m.Seconds = seconds
	m.Nanos = nanos

	return m.validateTimestamp()
}

// Value satisfies the valuer interface for database/sql.  Copied from ptypes.
func (m *Timestamp) Value() (driver.Value, error) {
	var t time.Time
	if m == nil {
		t = time.Unix(0, 0) // treat nil like the empty Timestamp
	} else {
		t = time.Unix(m.Seconds, int64(m.Nanos))
	}
	return t, m.validateTimestamp()
}

const (
	// Seconds field of the earliest valid Timestamp.
	// This is time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC).Unix().
	minValidSeconds = -62135596800
	// Seconds field just after the latest valid Timestamp.
	// This is time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC).Unix().
	maxValidSeconds = 253402300800
)

// validateTimestamp determines whether a Timestamp is valid.
// A valid timestamp represents a time in the range
// [0001-01-01, 10000-01-01) and has a Nanos field
// in the range [0, 1e9).
//
// If the Timestamp is valid, validateTimestamp returns nil.
// Otherwise, it returns an error that describes
// the problem.
//
// Every valid Timestamp can be represented by a time.Time, but the converse is not true.
func (m *Timestamp) validateTimestamp() error {
	if m.Seconds < minValidSeconds {
		return fmt.Errorf("timestamp: %v before 0001-01-01", m)
	}
	if m.Seconds >= maxValidSeconds {
		return fmt.Errorf("timestamp: %v after 10000-01-01", m)
	}
	if m.Nanos < 0 || m.Nanos >= 1e9 {
		return fmt.Errorf("timestamp: %v: nanos not in range [0, 1e9)", m)
	}
	return nil
}
