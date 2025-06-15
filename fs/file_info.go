package fs

import (
	gofs "io/fs"
	"os"
	"path/filepath"
	"time"
)

// This file defines the FileInfo interface and its implementation, providing
// a comprehensive abstraction for file metadata. It is part of a cross-platform
// file system utility package, designed to work seamlessly across different
// operating systems by leveraging platform-specific implementations where necessary.

// FileInfo is an interface that describes the file information.
// It extends the standard go fs.FileInfo interface with additional methods
// to retrieve file path, title, extension, and various timestamps.
type FileInfo interface {
	Name() string  // base name of the file (excluding the path)
	Path() string  // path to the file (excluding the base name)
	Abs() string   // absolute path to the file
	Title() string // title of the file
	Ext() string   // extension of the file
	Size() int64   // length in bytes for regular files; system-dependent for others
	IsDir() bool   // abbreviation for Mode().IsDir()

	CreationTime() time.Time   // creation time
	LastAccessTime() time.Time // last access time
	LastWriteTime() time.Time  // last write time
}

// fileInfo is a structure that contains information about a file.
// It implements the FileInfo interface and the go fs.FileInfo interface,
// providing a concrete representation of file metadata.
type fileInfo struct {
	name  string
	path  string
	abs   string
	title string
	ext   string
	size  int64
	mode  gofs.FileMode
	dir   bool

	creationTime   time.Time
	lastAccessTime time.Time
	lastWriteTime  time.Time
}

// fileInfo should implement the FileInfo interface
var _ FileInfo = (*fileInfo)(nil)

// fileInfo should implement the FileInfo interface
var _ gofs.FileInfo = (*fileInfo)(nil)

// NewFileInfo creates a new FileInfo struct.
// It takes a file path as input, retrieves the file's metadata using os.Stat,
// and returns a fileInfo object populated with this metadata.
func NewFileInfo(path string) (*fileInfo, error) {
	resolvedPath, err := Resolve(path)
	if err != nil {
		resolvedPath = path // Fallback to original path if symlink resolution fails
	}

	info, err := os.Stat(resolvedPath)
	if err != nil {
		return nil, err
	}
	return NewFileInfoFromFileInfo(info, filepath.Dir(resolvedPath))
}

// NewFileInfoFromFileInfo creates a new fileInfo struct from an existing os.FileInfo object.
// It extracts and computes various file attributes, such as name, path, size, and timestamps,
// and returns a fully populated fileInfo object.
func NewFileInfoFromFileInfo(info os.FileInfo, dir string) (*fileInfo, error) {
	f := fileInfo{}
	f.name = info.Name()
	f.path = filepath.Clean(dir)
	f.abs, _ = Resolve(f.path)
	f.ext = filepath.Ext(f.name)
	f.title = f.name[:len(f.name)-len(f.ext)]
	f.mode = info.Mode()
	f.dir = isDir(info)

	fullPath := filepath.Join(f.path, f.name)
	f.size = GetSize(info, fullPath)

	f.creationTime = getCreationTime(info)
	f.lastAccessTime = getLastAccessTime(info)
	f.lastWriteTime = getLastWriteTime(info)
	return &f, nil
}

// Name returns the base name of the file.
func (f fileInfo) Name() string {
	return f.name
}

// Path returns the path to the file.
func (f fileInfo) Path() string {
	return f.path
}

// Abs returns the absolute path to the file.
func (f fileInfo) Abs() string {
	return f.abs
}

// Title returns the title of the file.
func (f fileInfo) Title() string {
	return f.title
}

// Ext returns the extension of the file.
func (f fileInfo) Ext() string {
	return f.ext
}

// Size returns the length in bytes for regular files; system-dependent for others.
func (f fileInfo) Size() int64 {
	return f.size
}

// Mode returns the file mode bits.
func (f fileInfo) Mode() gofs.FileMode {
	return f.mode
}

// ModTime returns the modification time.
func (f fileInfo) ModTime() time.Time {
	return f.LastWriteTime()
}

// IsDir returns true if the file is a directory.
func (f fileInfo) IsDir() bool {
	return f.dir
}

// Sys returns the underlying data source.
// This is a placeholder implementation and can be modified to return
// actual underlying data if needed.
func (f fileInfo) Sys() any {
	return nil
}

// CreationTime returns the creation time.
func (f fileInfo) CreationTime() time.Time {
	return f.creationTime
}

// LastAccessTime returns the last access time.
func (f fileInfo) LastAccessTime() time.Time {
	return f.lastAccessTime
}

// LastWriteTime returns the last write time.
func (f fileInfo) LastWriteTime() time.Time {
	return f.lastWriteTime
}
