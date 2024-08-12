package file

import (
	"fmt"
	"os"
	"path/filepath"
)

// FilePathWalkDir walks through the directory tree rooted at 'root' and returns a slice of file paths. It skips
// directories and returns an error if any issue is encountered during the walk.
func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return fmt.Errorf("nil FileInfo for path: %s", path)
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error walking the directory tree: %v", err)
	}
	return files, nil
}
