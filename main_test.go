package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/lyckade/gositm/timestamp"
)

func TestMakeFilename(t *testing.T) {
	filename := "a/b/MyFile.txt"
	lastChange := time.Now()
	ts := timestamp.FromTime(lastChange)
	expect := fmt.Sprintf("a/b/MyFile.%s.txt", ts)
	got := MakeBackupFilename(filename, lastChange)
	if expect != got {
		t.Fatalf("Expect: %v Got: %v", expect, got)
	}
}

func TestMakeBackupDir(t *testing.T) {
	filepath := "rootDir/a/b/c/myFile.txt"
	expect := "rootDir/backup/a\\b\\c\\myFile.txt"
	got, err := MakeBackupDir(filepath, "rootDir", "backup")
	if err != nil {
		t.Error(err)
	}
	if expect != got {
		t.Fatalf("Expect: %v Got: %v", expect, got)
	}
}

func TestIgnorePath(t *testing.T) {
	ignorePatterns := []string{".git", "test"}
	paths := []struct {
		p      string
		expect bool
	}{
		{
			"123/23123/",
			false,
		},
		{
			".git/here",
			true,
		},
		{
			"test/1234",
			true,
		},
		{
			"myFile",
			false,
		},
	}
	for _, testCase := range paths {
		got, _ := MatchPath(testCase.p, ignorePatterns)
		if testCase.expect != got {
			t.Fatalf("Path: %v\nExpect: %v Got: %v",
				testCase.p,
				testCase.expect,
				got)
		}
	}

}
