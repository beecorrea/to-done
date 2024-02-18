package todo

import (
	"os"
)

// Reads a file and returns its content along with an error, if any.
func GetFileContents(path string) (string, error) {
	c, err := os.ReadFile(path)
	return string(c), err
}
