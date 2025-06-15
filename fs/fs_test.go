package fs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFile(t *testing.T) {
	testCases := []struct {
		name     string
		path     string
		expected bool
	}{
		{"existing file", filepath.Join("testdata", "directory", "text.txt"), true},
		{"existing directory", filepath.Join("testdata", "directory"), false},
		{"symlink to file", filepath.Join("testdata", "linked", "text.txt"), true},
		{"symlink to directory", filepath.Join("testdata", "linked"), false},
		{"non-existing path", "testdata/nonexistent.txt", false},
		{"empty path", "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// On Windows, creating symlinks requires special permissions.
			// The setup function creates a fallback directory if symlink creation fails.
			// Skip symlink-related tests if the symlink was not created.
			if (tc.name == "symlink to file" || tc.name == "symlink to directory") && !symlinkCreated {
				t.Skip("Skipping symlink test because symlink could not be created")
			}
			assert.Equal(t, tc.expected, IsFile(tc.path))
		})
	}
}

func TestIsDir(t *testing.T) {
	testCases := []struct {
		name     string
		path     string
		expected bool
	}{
		{"existing file", filepath.Join("testdata", "directory", "text.txt"), false},
		{"existing directory", filepath.Join("testdata", "directory"), true},
		{"symlink to directory", filepath.Join("testdata", "linked"), true},
		{"symlink to file", filepath.Join("testdata", "linked", "text.txt"), false},
		{"non-existing path", "testdata/nonexistent", false},
		{"empty path", "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if (tc.name == "symlink to file" || tc.name == "symlink to directory") && !symlinkCreated {
				t.Skip("Skipping symlink test because symlink could not be created")
			}
			assert.Equal(t, tc.expected, IsDir(tc.path))
		})
	}
}

func TestIsEmpty(t *testing.T) {
	testCases := []struct {
		name     string
		path     string
		expected bool
	}{
		{"empty file", filepath.Join("testdata", "empty_file.txt"), true},
		{"non-empty file", filepath.Join("testdata", "directory", "text.txt"), false},
		{"empty directory", filepath.Join("testdata", "empty_dir"), true},
		{"non-empty directory", filepath.Join("testdata", "directory"), false},
		{"non-existing path", "testdata/nonexistent.txt", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, IsEmpty(tc.path))
		})
	}
}

func TestResolve(t *testing.T) {
	// Test case for a simple file path
	t.Run("simple path", func(t *testing.T) {
		path := filepath.Join("testdata", "directory", "text.txt")
		absPath, err := filepath.Abs(path)
		assert.NoError(t, err)

		resolvedPath, err := Resolve(path)
		assert.NoError(t, err)
		assert.Equal(t, absPath, resolvedPath)
	})

	// Test case for a path that does not exist
	t.Run("non-existing path", func(t *testing.T) {
		path := "testdata/nonexistent.txt"
		_, err := Resolve(path)
		assert.Error(t, err)
		assert.True(t, os.IsNotExist(err))
	})

	// Test case for a symlink
	t.Run("symlink path", func(t *testing.T) {
		if !symlinkCreated {
			t.Skip("Skipping symlink test because symlink could not be created")
		}
		path := filepath.Join("testdata", "linked")
		targetPath := filepath.Join("testdata", "directory")
		absTargetPath, err := filepath.Abs(targetPath)
		assert.NoError(t, err)

		resolvedPath, err := Resolve(path)
		assert.NoError(t, err)
		assert.Equal(t, absTargetPath, resolvedPath)
	})
}
