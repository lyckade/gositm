package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"time"

	"github.com/lyckade/gositm/conf"
	"github.com/lyckade/gositm/timestamp"
)

var properties = conf.Properties

var rootDir string

func init() {
	dir := filepath.Dir(os.Args[0])

	rootDir = dir
}

func main() {
	filepath.Walk(rootDir, func(fpath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return err
		}
		ignore, err := MatchPath(fpath, properties.Ignore)
		if err != nil || ignore {
			fmt.Println(fpath)
			return err
		}

		targetPath, err := MakeBackupDir(
			MakeBackupFilename(fpath, info.ModTime()),
			rootDir,
			properties.BackupFolder)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("In: %v\nOut:%v\n\n", fpath, targetPath)
		err = CopyFile(fpath, targetPath)
		if err != nil {
			log.Fatal(err)
		}
		return err
	})
}

// MakeBackupDir creates from a given path the path to that file inside
// of the Backup folder
func MakeBackupDir(fpath, rootDir, backupFolder string) (string, error) {
	bf := path.Join(rootDir, backupFolder)
	relPath, err := filepath.Rel(rootDir, fpath)
	return path.Join(bf, relPath), err
}

// MakeBackupFilename takes a filename and adds the timestamp before
// the extension.
// It is also possible to use a filepath as filename
func MakeBackupFilename(filename string, lastChange time.Time) string {
	ext := path.Ext(filename)
	filebase := filename[:len(filename)-len(ext)]
	if len(ext) == len(filename) {
		ext = ""
		filebase = filename
	}

	ts := timestamp.FromTime(lastChange)
	return fmt.Sprintf("%s.%s%s", filebase, ts, ext)
}

// MatchPath checks if a filepath matches a [] of patterns
func MatchPath(fpath string, patterns []string) (bool, error) {
	for _, p := range patterns {
		match, err := regexp.MatchString(p, fpath)
		if err != nil {
			return false, err
		}
		if match {
			return true, nil
		}
	}
	return false, nil
}

// CopyFile ensures that the target path exists and copies a file
func CopyFile(src, dest string) error {
	err := os.MkdirAll(filepath.Dir(dest), 0777)
	if err != nil {
		return err
	}
	fdest, err := os.Create(dest)
	defer fdest.Close()
	if err != nil {
		return err
	}

	fsrc, err := os.Open(src)
	defer fsrc.Close()
	if err != nil {
		return err
	}

	_, err = io.Copy(fdest, fsrc)
	if err != nil {
		return err
	}
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	err = os.Chtimes(dest, srcInfo.ModTime(), srcInfo.ModTime())
	if err != nil {
		return err
	}
	return nil
}
