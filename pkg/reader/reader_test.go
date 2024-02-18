package reader

import (
	"fmt"
	"strings"
	"testing"
)

func TestReadFile(t *testing.T) {
	// todo: refactor this to be passed as args from cli.
	file := "../../data/test.go"
	if _, err := GetFileContents(file); err != nil {
		t.Fatalf("should read file, but got error %s", err.Error())
	}
}

func TestReadDir(t *testing.T) {
	dir := "../../data"
	_, err := GetFilesInDir(dir)
	if err != nil {
		t.Fatalf("should read all files in dir.")
	}
}

func TestReadDirRecursive(t *testing.T) {
	dir := parentDir
	_, err := GetFilesInDir(dir)
	if err != nil {
		t.Fatalf("should read all files in child dirs.")
	}
}

func TestReadDirCurrent(t *testing.T) {
	dir := currentDir
	_, err := GetFilesInDir(dir)
	if err != nil {
		t.Fatalf("should read all files in current dir.")
	}
}

func TestReadDirCurrentSlash(t *testing.T) {
	dir := currentDirSlash
	_, err := GetFilesInDir(dir)
	if err != nil {
		t.Fatalf("should read all files in current dir even if dir has slash.")
	}
}

func TestReadDirIgnore(t *testing.T) {
	dir := "../../"
	files, err := GetFilesInDir(dir)
	if err != nil {
		t.Fatalf("should read all files in dir '%s'", dir)
	}
	for _, f := range files {
		fmt.Println(f)
		for k := range blocklist {
			if strings.Contains(f, k) {
				t.Fatalf("should ignore file if it's in blocklist.")
			}
		}
	}
}
