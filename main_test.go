package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/lyckade/gositm/timestamp"
)

func TestMakeFilename(t *testing.T) {
	filename := "MyFile.txt"
	lastChange := time.Now()
	ts := timestamp.FromTime(lastChange)
	expect := fmt.Sprintf("MyFile.%s.txt", ts)
	got := MakeFilename(filename, lastChange)
	if expect != got {
		t.Fatalf("Expect: %v Got: %v", expect, got)
	}
}
