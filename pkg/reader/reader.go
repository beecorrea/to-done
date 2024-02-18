package reader

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	currentDir      = "."
	currentDirSlash = "./"
	parentDir       = "../"
)

var blocklist = map[string]struct{}{".git": {}}

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
		// todo: this should be a dependency.
		if _, ok := blocklist[entry.Name()]; !ok {
			fullPath := filepath.Join(absPath, entry.Name())
			// todo: refactor to another function. maybe GetFilesFrom(path).
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
	}

	return pathsList, nil
}

// Checks if the path has '.', './' or '../'.
func hasSpecialReference(path string) bool {
	isCurrDir := path == currentDir || strings.HasPrefix(path, currentDirSlash)
	isParentDir := strings.HasPrefix(path, parentDir)
	return isCurrDir || isParentDir
}

// If path is '.' or './', get absolute path to current dir.
func resolvePath(path string) string {
	needsResolution := hasSpecialReference(path)
	if needsResolution {
		p, err := filepath.Abs(path)
		if err != nil {
			panic(err)
		}

		return p
	}

	return path
}

// Removes a root from the path so it won't show up in pretty-print.
func StripRootFromPath(root string, path string) string {
	abs := resolvePath(root)
	trailing := fmt.Sprintf("%s/", abs)
	return strings.Replace(path, trailing, "", -1)
}

type WalkFn func(string) string
type LineInfo struct {
	Content      string
	Number       int
	Filename     string
	PathFromRoot string
}

// Applies a WalkFn throughout a file and generates LineInfo for every successful apply.
func WalkFile(path string, fn WalkFn) ([]LineInfo, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	infos := make([]LineInfo, 0)
	scanner := bufio.NewScanner(f)
	currLine := 1

	for scanner.Scan() {
		line := scanner.Text()
		apply := fn(line)
		if apply != "" {
			lineInfo := LineInfo{Content: apply, Number: currLine, Filename: path}
			infos = append(infos, lineInfo)
		}
		currLine++
	}

	return infos, nil
}

// Auxiliary function to group by filename.
func GroupByFilename(lines []LineInfo) map[string][]LineInfo {
	groups := make(map[string][]LineInfo, 0)
	for _, line := range lines {
		groups[line.Filename] = append(groups[line.Filename], line)
	}
	return groups
}
