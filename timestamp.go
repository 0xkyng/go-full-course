package main

import (
	"strconv"
	"time"
)

type timestamp struct {
	time.Time
}

func (ts timestamp) String() string {
	if ts.IsZero() {
		return "Unknown"
	}

	const layout = "2006/01"
	return ts.Format(layout)
}

// Unmarshaler
func (ts *timestamp) UnmarshalJSON(data []byte) error {
	*ts = toTimestamp(string(data))
	return nil
}

// Marshaler
func (ts timestamp) MarshalJSON() (data []byte, _ error) {
	// ts -> integer -> ts.Unix() -> integer
	// data <- interger -> strconv.AppendInt(data, integer, 10)
	return strconv.AppendInt(data, ts.Unix(), 10), nil


}

func toTimestamp(v interface{}) (ts timestamp) {
	var t int

	switch v := v.(type) {
	case int:
		// book{title: "moby dick", price: 10, published: 118281600},
		t = v
	case string:
		// book{title: "odyssey", price: 15, published: "733622400"},
		t, _ = strconv.Atoi(v)
	
	}

	

	ts.Time = time.Unix(int64(t), 0)
	return ts
}