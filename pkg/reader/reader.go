package reader

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	CurrentDir      = "."
	CurrentDirSlash = "./"
	ParentDir       = "../"
)

// Reads a file and returns its content along with an error, if any.
func GetFileContents(path string) (string, error) {
	c, err := os.ReadFile(path)
	return string(c), err
}

// Reads a directory and returns all files in that dir.
// If a file is a directory, the function recurses to that directory.
func GetFilesInDir(path string) ([]string, error) {
	absPath := resolvePath(path)
	dir, err := os.ReadDir(absPath)
	if err != nil {
		return nil, err
	}

	pathsList := make([]string, 0)
	for _, entry := range dir {
		fullPath := filepath.Join(absPath, entry.Name())

		if !entry.IsDir() {
			pathsList = append(pathsList, fullPath)
		} else {
			paths, err := GetFilesInDir(fullPath)
			if err != nil {
				return nil, err
			}
			pathsList = append(pathsList, paths...)
		}
	}

	return pathsList, nil
}

// If path is '.' or './', get absolute path to current dir.
func resolvePath(path string) string {
	isCurrDir := path == CurrentDir || strings.HasPrefix(path, CurrentDirSlash)
	isParentDir := strings.HasPrefix(path, ParentDir)
	needsResolution := isCurrDir || isParentDir

	if needsResolution {
		p, err := filepath.Abs(path)
		if err != nil {
			panic(err)
		}

		return p
	}

	return path
}
