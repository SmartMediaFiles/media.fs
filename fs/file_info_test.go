package fs

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// symlinkCreated is a package-level variable that tracks whether the
// symbolic link was successfully created during test setup. This is used
// to conditionally skip tests that depend on symlinks if the environment
// (e.g., Windows without admin rights) does not support their creation.
var symlinkCreated bool

// setupTestData creates a temporary directory structure and files for testing.
// It sets up a 'testdata' directory containing:
// 1. A directory named 'directory' with a 'text.txt' file inside.
// 2. A symbolic link named 'linked' pointing to the 'directory'.
// If creating the symlink fails (e.g., due to lack of privileges on Windows),
// it creates a regular directory named 'linked' and copies the content from
// 'directory' into it as a fallback.
func setupTestData() {
	// Create the root testdata directory if it doesn't exist
	testDataDir := "testdata"
	if err := os.MkdirAll(testDataDir, 0755); err != nil {
		panic("Failed to create testdata directory: " + err.Error())
	}

	// Create an empty directory
	emptyDirPath := filepath.Join(testDataDir, "empty_dir")
	if err := os.MkdirAll(emptyDirPath, 0755); err != nil {
		panic("Failed to create 'empty_dir': " + err.Error())
	}

	// Create an empty file
	emptyFilePath := filepath.Join(testDataDir, "empty_file.txt")
	if err := os.WriteFile(emptyFilePath, []byte{}, 0644); err != nil {
		panic("Failed to create empty_file.txt: " + err.Error())
	}

	// Create the 'directory' subdirectory
	directoryPath := filepath.Join(testDataDir, "directory")
	if err := os.MkdirAll(directoryPath, 0755); err != nil {
		panic("Failed to create 'directory' subdirectory: " + err.Error())
	}

	// Create 'text.txt' in the 'directory' subdirectory
	textFilePath := filepath.Join(directoryPath, "text.txt")
	textContent := "Some text file content.\n" // 24 bytes
	if err := os.WriteFile(textFilePath, []byte(textContent), 0644); err != nil {
		panic("Failed to create text.txt: " + err.Error())
	}

	// Attempt to create a symbolic link.
	linkedPath := filepath.Join(testDataDir, "linked")

	// On Windows, it's safer to use an absolute path for the symlink target.
	absDirectoryPath, err := filepath.Abs(directoryPath)
	if err != nil {
		panic("Failed to get absolute path for symlink target: " + err.Error())
	}

	if err := os.Symlink(absDirectoryPath, linkedPath); err != nil {
		// If creating the symlink fails, set the flag to false.
		// Then, create a regular directory as a fallback and copy the content.
		// This is a common scenario on Windows without administrator privileges.
		symlinkCreated = false
		// Create the 'linked' directory.
		if err := os.MkdirAll(linkedPath, 0755); err != nil {
			panic("Failed to create fallback 'linked' directory: " + err.Error())
		}
		// Copy the content from 'directory' to 'linked'.
		if err := copyDir(directoryPath, linkedPath); err != nil {
			panic("Failed to copy content to fallback 'linked' directory: " + err.Error())
		}
	} else {
		// If the symlink is created successfully, set the flag to true.
		symlinkCreated = true
	}
}

// copyDir recursively copies a directory from src to dst.
func copyDir(src, dst string) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := os.MkdirAll(dstPath, 0755); err != nil {
				return err
			}
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}
	return nil
}

// copyFile copies a single file from src to dst.
func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}

// cleanupTestData removes the temporary 'testdata' directory and all its
// contents. It's called after all tests in the package have run.
func cleanupTestData() {
	// Clean up the testdata directory
	testDataDir := "testdata"
	if err := os.RemoveAll(testDataDir); err != nil {
		panic("Failed to clean up testdata directory: " + err.Error())
	}
}

// TestMain is a special test function that serves as the entry point for
// running tests in this package. It handles the setup and teardown logic
// by calling cleanupTestData after all other tests have completed.
// The setup is handled by the init() function, which runs before TestMain.
func TestMain(m *testing.M) {
	// Create the test data structure
	setupTestData()

	// Run all tests in the package.
	exitCode := m.Run()

	// Clean up test resources after all tests have run.
	cleanupTestData()

	// Exit with the appropriate status code.
	os.Exit(exitCode)
}

// TestNewFileInfo contains a suite of sub-tests for the NewFileInfo function,
// covering various file types and structures like directories, files, and symlinks.
func TestNewFileInfo(t *testing.T) {
	// Test case for a standard directory.
	t.Run("directory", func(t *testing.T) {
		path := filepath.Join("testdata", "directory")

		info, err := NewFileInfo(path)
		if err != nil {
			t.Fatal(err)
		}

		resolvedPath, err := Resolve("testdata")
		if err != nil {
			t.Fatal("Failed to resolve symlink:", err)
		}

		assert.Equal(t, "directory", info.Name())
		assert.Equal(t, resolvedPath, info.Path())
		assert.Equal(t, "directory", info.Title())
		assert.Equal(t, "", info.Ext())
		assert.NotEqual(t, int64(0), info.Size())
		assert.Equal(t, true, info.IsDir())
	})

	// Test case for a standard text file within a directory.
	t.Run("directory_text.txt", func(t *testing.T) {
		path := filepath.Join("testdata", "directory", "text.txt")

		info, err := NewFileInfo(path)
		if err != nil {
			t.Fatal(err)
		}

		resolvedPath, err := Resolve(filepath.Join("testdata", "directory"))
		if err != nil {
			t.Fatal("Failed to resolve symlink:", err)
		}

		assert.Equal(t, "text.txt", info.Name())
		assert.Equal(t, resolvedPath, info.Path())
		assert.Equal(t, "text", info.Title())
		assert.Equal(t, ".txt", info.Ext())
		assert.Equal(t, int64(24), info.Size())
		assert.Equal(t, false, info.IsDir())
	})

	// Test case for a symbolic link or its fallback directory copy.
	// This test works whether a symlink or a regular directory was created.
	t.Run("linked", func(t *testing.T) {
		path := filepath.Join("testdata", "linked")

		info, err := NewFileInfo(path)
		if err != nil {
			t.Fatal(err)
		}

		resolvedPath, err := Resolve("testdata")
		if err != nil {
			t.Fatal("Failed to resolve symlink:", err)
		}

		expectedName := "linked"
		if symlinkCreated {
			expectedName = "directory" // If symlink was created, the name is 'linked'.
		}

		assert.Equal(t, expectedName, info.Name())
		assert.Equal(t, resolvedPath, info.Path())
		assert.Equal(t, expectedName, info.Title())
		assert.Equal(t, "", info.Ext())
		assert.NotEqual(t, int64(0), info.Size())
		assert.Equal(t, true, info.IsDir())
	})

	// Test case for a text file inside the linked/copied directory.
	t.Run("linked_text.txt", func(t *testing.T) {
		path := filepath.Join("testdata", "linked", "text.txt")

		info, err := NewFileInfo(path)
		if err != nil {
			t.Fatal(err)
		}

		resolvedPath, err := Resolve(filepath.Join("testdata", "linked"))
		if err != nil {
			t.Fatal("Failed to resolve symlink:", err)
		}

		assert.Equal(t, "text.txt", info.Name())
		assert.Equal(t, resolvedPath, info.Path())
		assert.Equal(t, "text", info.Title())
		assert.Equal(t, ".txt", info.Ext())
		assert.Equal(t, int64(24), info.Size())
		assert.Equal(t, false, info.IsDir())
	})
}
