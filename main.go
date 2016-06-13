package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/lyckade/gositm/conf"
	"github.com/lyckade/gositm/timestamp"
)

var properties = conf.Properties

var rootDir string

func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	rootDir = dir
}

func main() {

}

// MakeFilename takes a filename and adds the timestamp before
// the extension.
// It is also possible to use a filepath as filename
func MakeFilename(filename string, lastChange time.Time) string {
	ext := path.Ext(filename)
	ts := timestamp.FromTime(lastChange)
	return fmt.Sprintf("%s.%s%s", filename[:len(filename)-len(ext)], ts, ext)
}
