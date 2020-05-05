package ptypes

import (
	tspb "github.com/golang/protobuf/ptypes/timestamp"
	"testing"
)

func TestTimestampScan(t *testing.T) {
	tests := []struct {
		name string
		value   string
		seconds int64
		wantErr bool
	}{
		{
			name: "6a8f72bd-7d1c-4949-a68a-6a2b31093b8a - scan RFC3339",
			value: "2002-10-02T10:00:00-05:00",
			seconds: 1033570800,
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


func TestTimestamp_IsZero(t *testing.T) {
	tests := []struct {
		name   string
		ts *tspb.Timestamp
		want   bool
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
			ts: &tspb.Timestamp{},
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
