package todo

import "testing"

func TestReadFile(t *testing.T) {
	// todo: refactor this to be passed as args from cli.
	file := "../../data/test.go"
	if _, err := GetFileContents(file); err != nil {
		t.Fatalf("should read file, but got error %s", err.Error())
	}
}
