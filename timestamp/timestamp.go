package timestamp

import "time"

//Timestamp represents a date time
type Timestamp string

//TimestampLayout represents the Layout of the timestamp
var TimestampLayout = "20060102150405"

// Parse generates a timestamp from a string and returns the pointer.
// The layoutString has to be in the format of the time package of
// the standardlib
func Parse(layoutString string, s string) *Timestamp {
	t, _ := time.Parse(layoutString, s)
	return FromTime(t)
}

// FromTime generates a timestamp from a time.Time object and returns
// a pointer.
func FromTime(t time.Time) *Timestamp {
	ts := Timestamp(t.Format(TimestampLayout))
	return &ts
}

// String implements the stringer interface
func (t *Timestamp) String() string {
	return string(*t)
}
