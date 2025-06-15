package fs

import (
	"os"
	"path/filepath"
	"strings"
)

// IsFile returns true if file exists and is not a directory.
func IsFile(fileName string) bool {
	if fileName == "" {
		return false
	}

	info, err := os.Stat(fileName)

	return err == nil && !isDir(info)
}

// IsDir returns true if file exists and is a directory.
func IsDir(fileName string) bool {
	if fileName == "" {
		return false
	}

	info, err := os.Stat(fileName)

	return err == nil && isDir(info)
}

// isDir returns true if file exists and is a directory or symlink.
func isDir(fileInfo os.FileInfo) bool {
	if fileInfo == nil {
		return false
	}

	m := fileInfo.Mode()

	return m&os.ModeDir != 0 || m&os.ModeSymlink != 0
}

// IsEmpty returns true if the destination is empty.
// If the destination is a file, it returns true if the file exists and is empty.
// If the destination is a directory, it returns true if the directory exists and is empty.
func IsEmpty(destination string) bool {
	if IsFile(destination) {
		return isEmptyFile(destination)
	} else if IsDir(destination) {
		return isEmptyDir(destination)
	} else {
		return false
	}
}

// isEmptyFile returns true if the file exists and is empty.
func isEmptyFile(fileName string) bool {
	info, err := os.Stat(fileName)
	return err == nil && info.Size() == 0
}

// isEmptyDir returns true if the directory exists and is empty.
func isEmptyDir(dirName string) bool {
	dir, err := os.ReadDir(dirName)
	return err == nil && len(dir) == 0
}

// Resolve returns the absolute path to the file.
// If the file path contains a tilde, it will be expanded to the home directory.
// If the file path contains symbolic links, they will be resolved.
func Resolve(filePath string) (string, error) {
	if filePath == "" {
		return "", os.ErrNotExist
	}

	// Expand tilde to home directory
	if strings.HasPrefix(filePath, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		filePath = filepath.Join(home, filePath[1:])
	}

	// Resolve symbolic links and get absolute path
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return "", err
	}

	resolvedPath, err := filepath.EvalSymlinks(absPath)
	if err != nil {
		return "", err
	}

	// Check if the file exists
	if _, err := os.Stat(resolvedPath); err != nil {
		return "", err
	}

	return resolvedPath, nil
}
