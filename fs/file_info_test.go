package fs

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestNewFileInfo(t *testing.T) {
	t.Run("directory", func(t *testing.T) {
		info, err := NewFileInfo().FromPath(filepath.Join("testdata", "directory"))
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "directory", info.Name())
		assert.Equal(t, "testdata", info.Path())
		assert.Equal(t, "directory", info.Title())
		assert.Equal(t, "", info.Ext())
		assert.Equal(t, int64(0), info.Size())
		assert.Equal(t, true, info.IsDir())
	})

	t.Run("directory_text.txt", func(t *testing.T) {
		info, err := NewFileInfo().FromPath(filepath.Join("testdata", "directory", "text.txt"))
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "text.txt", info.Name())
		assert.Equal(t, filepath.Join("testdata", "directory"), info.Path())
		assert.Equal(t, "text", info.Title())
		assert.Equal(t, ".txt", info.Ext())
		assert.Equal(t, int64(21), info.Size())
		assert.Equal(t, false, info.IsDir())
	})

	t.Run("linked", func(t *testing.T) {
		i, err := os.Stat(filepath.Join("testdata", "linked"))
		if err != nil {
			t.Fatal(err)
		}

		info, err := NewFileInfo().FromFileInfo(i, "testdata")
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "linked", info.Name())
		assert.Equal(t, "testdata", info.Path())
		assert.Equal(t, "linked", info.Title())
		assert.Equal(t, "", info.Ext())
		assert.Equal(t, int64(0), info.Size())
		assert.Equal(t, true, info.IsDir())
	})

	t.Run("linked_text.txt", func(t *testing.T) {
		info, err := NewFileInfo().FromPath(filepath.Join("testdata", "linked", "text.txt"))
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "text.txt", info.Name())
		assert.Equal(t, filepath.Join("testdata", "linked"), info.Path())
		assert.Equal(t, "text", info.Title())
		assert.Equal(t, ".txt", info.Ext())
		assert.Equal(t, int64(21), info.Size())
		assert.Equal(t, false, info.IsDir())
	})

	t.Run("linked_directory", func(t *testing.T) {
		info, err := NewFileInfo().FromPath(filepath.Join("testdata", "linked", "directory"))
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "directory", info.Name())
		assert.Equal(t, filepath.Join("testdata", "linked"), info.Path())
		assert.Equal(t, "directory", info.Title())
		assert.Equal(t, "", info.Ext())
		assert.Equal(t, int64(0), info.Size())
		assert.Equal(t, true, info.IsDir())
	})

	t.Run("linked_directory_text.txt", func(t *testing.T) {
		info, err := NewFileInfo().FromPath(filepath.Join("testdata", "linked", "directory", "text.txt"))
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "text.txt", info.Name())
		assert.Equal(t, filepath.Join("testdata", "linked", "directory"), info.Path())
		assert.Equal(t, "text", info.Title())
		assert.Equal(t, ".txt", info.Ext())
		assert.Equal(t, int64(21), info.Size())
		assert.Equal(t, false, info.IsDir())
	})
}
