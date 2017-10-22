package timestamp

import (
	"errors"
	"time"

	"database/sql/driver"

	"github.com/golang/protobuf/ptypes"
)

func (m *Timestamp) Scan(value interface{}) error {
	if myTime, ok := value.(time.Time); ok {
		foo, err := ptypes.TimestampProto(myTime)
		if err != nil {
			return err
		}
		m = foo
		return nil
	}
	return errors.New("incompatible type passed, expected time.Time")
}

func (m *Timestamp) Value() (driver.Value, error) {
	return ptypes.Timestamp(m)
}
