package reader

import (
	"fmt"
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
	files, err := GetFilesInDir(dir)
	fmt.Println(files)
	if err != nil {
		t.Fatalf("should read all files in dir.")
	}
}

func TestReadDirRecursive(t *testing.T) {
	dir := "../"
	files, err := GetFilesInDir(dir)
	fmt.Println(files)
	if err != nil {
		t.Fatalf("should read all files in child dirs.")
	}
}

func TestReadDirCurrent(t *testing.T) {
	dir := CurrentDir
	files, err := GetFilesInDir(dir)
	fmt.Println(files)
	if err != nil {
		t.Fatalf("should read all files in current dir.")
	}
}

func TestReadDirCurrentSlash(t *testing.T) {
	dir := CurrentDirSlash
	files, err := GetFilesInDir(dir)
	fmt.Println(files)
	if err != nil {
		t.Fatalf("should read all files in current dir even if dir has slash.")
	}
}
