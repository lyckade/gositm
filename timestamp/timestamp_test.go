package timestamp

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	layout := "2006.01.02 15:04:05"
	expect := Timestamp("20160124115500")
	got := Parse(layout, "2016.01.24 11:55:00")
	if expect != *got {
		t.Fatalf("Expected: %v Got: %v", expect, *got)
	}
}

func TestFromTime(t *testing.T) {
	layout := "2006.01.02 15:04:05"
	inputTime, _ := time.Parse(layout, "2016.01.24 11:55:00")
	expect := Timestamp("20160124115500")
	got := FromTime(inputTime)
	if expect != *got {
		t.Fatalf("Expected: %v Got: %v", expect, *got)
	}
}

func TestString(t *testing.T) {
	ts := Parse("2006.01.02 15:04:05", "2016.01.24 11:55:00")
	expect := "20160124115500"
	if ts.String() != expect {
		t.Fatalf("Expected: %v Got: %v", expect, ts.String())
	}
}
