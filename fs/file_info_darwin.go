//go:build darwin

package fs

import (
	"os"
	"syscall"
	"time"
)

// This file provides Darwin-specific (macOS) implementations for retrieving
// file timestamps. It is part of a cross-platform file system utility package.

// getCreationTime returns the creation time (birth time) of a file.
// On Darwin, this is accessed via the `Birthtimespec` field of `syscall.Stat_t`.
func getCreationTime(info os.FileInfo) time.Time {
	stat := info.Sys().(*syscall.Stat_t)
	return time.Unix(stat.Birthtimespec.Sec, stat.Birthtimespec.Nsec)
}

// getLastAccessTime retrieves the last access time of a file.
// It uses the `Atimespec` field from the `syscall.Stat_t` structure.
func getLastAccessTime(info os.FileInfo) time.Time {
	stat := info.Sys().(*syscall.Stat_t)
	return time.Unix(stat.Atimespec.Sec, stat.Atimespec.Nsec)
}

// getLastWriteTime retrieves the last modification time of a file.
// It uses the `Mtimespec` field from the `syscall.Stat_t` structure.
func getLastWriteTime(info os.FileInfo) time.Time {
	stat := info.Sys().(*syscall.Stat_t)
	return time.Unix(stat.Mtimespec.Sec, stat.Mtimespec.Nsec)
}

// GetSize returns the size of a file or directory. On Darwin, the size
// reported by `os.FileInfo` is accurate for both files and directories,
// so this function simply returns that value.
func GetSize(info os.FileInfo, path string) int64 {
	return info.Size()
}
