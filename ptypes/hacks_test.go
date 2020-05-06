package ptypes

import (
	"testing"
	"time"

	tspb "github.com/golang/protobuf/ptypes/timestamp"
)

func TestTimestampScan(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		seconds int64
		wantErr bool
	}{
		{
			name:    "6a8f72bd-7d1c-4949-a68a-6a2b31093b8a - scan RFC3339",
			value:   "1969-12-31T19:00:00-05:00",
			seconds: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &tspb.Timestamp{}
			err := ts.Scan(tt.value)

			//error
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if tt.wantErr {
				return //nothing else to test
			}

			if tt.seconds != ts.Seconds {
				t.Errorf("value not scanned correctly, wanted: %d, got: %d", tt.seconds, ts.Seconds)
				t.Fail()
			}
		})
	}
}

func TestTimestampValue(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		seconds int64
		wantErr bool
	}{
		{
			name:    "cf0d9182-7b29-43a3-b4b2-648b1d30ab94",
			value:   "2002-10-02T11:00:00-04:00",
			seconds: 1033570800,
		},
		{
			name:    "488d766f-d5e2-485f-b9ce-79ce93ee3191",
			value:   "2002-11-17T17:06:40-05:00",
			seconds: 1037570800,
		},
		{
			name:    "9a432c7b-16a6-4da7-95bf-e7693123a9ab",
			value:   "2006-01-18T02:53:20-05:00",
			seconds: 1137570800,
		},
		{
			name:    "6c689644-1dbc-41f6-a127-ffb568fe16cb",
			value:   "1969-12-31T19:00:00-05:00",
			seconds: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &tspb.Timestamp{Seconds: tt.seconds}
			val, err := ts.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if tt.wantErr {
				return //nothing else to test
			}

			if tval, ok := val.(time.Time); ok {
				//What is this
				formatted := tval.Format("2006-01-02T15:04:05-07:00")
				if formatted != tt.value {
					t.Errorf("wanted: %s, got %s", tt.value, formatted)
					t.Fail()
				}
			}

		})
	}
}

func TestTimestampValueNil(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		seconds int64
		wantErr bool
	}{
		{
			name:    "cf0d9182-7b29-43a3-b4b2-648b1d30ab94",
			value:   "2002-10-02T11:00:00-04:00",
			seconds: 1033570800,
		},
		{
			name:    "488d766f-d5e2-485f-b9ce-79ce93ee3191",
			value:   "2002-11-17T17:06:40-05:00",
			seconds: 1037570800,
		},
		{
			name:    "9a432c7b-16a6-4da7-95bf-e7693123a9ab",
			value:   "2006-01-18T02:53:20-05:00",
			seconds: 1137570800,
		},
		{
			name:    "6c689644-1dbc-41f6-a127-ffb568fe16cb",
			value:   "1969-12-31T19:00:00-05:00",
			seconds: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ts *tspb.Timestamp
			val, err := ts.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if tt.wantErr {
				return //nothing else to test
			}

			if tval, ok := val.(time.Time); ok {
				//What is this
				formatted := tval.Format("2006-01-02T15:04:05-07:00")
				if formatted != tt.value {
					t.Errorf("wanted: %s, got %s", tt.value, formatted)
					t.Fail()
				}
			}

		})
	}
}

func TestTimestamp_IsZero(t *testing.T) {
	tests := []struct {
		name string
		ts   *tspb.Timestamp
		want bool
	}{
		{
			name: "aa56851e-98b8-47e6-b0ef-f7a1a053666a",
			want: true,
		},
		{
			name: "c0e4cef6-503b-46d5-8a6d-1d8376b0c1a2",
			ts: &tspb.Timestamp{
				Seconds: 1,
			},
			want: false,
		},
		{
			name: "2a946b11-5b63-408d-b91d-34a4e3e80e69",
			ts:   &tspb.Timestamp{},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ts.IsZero(); got != tt.want {
				t.Errorf("Timestamp.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
