//go:build linux || freebsd || netbsd || openbsd

package fs

import (
	"os"
	"syscall"
	"time"
)

// This file provides Unix-specific implementations for retrieving file timestamps.
// It is part of a cross-platform file system utility package, where different
// implementations are provided for different operating systems using build tags.

// getCreationTime returns the creation time of a file.
// It uses the `Ctim` field from `syscall.Stat_t`, which on Linux represents
// the last status change time and is the closest equivalent to creation time.
func getCreationTime(info os.FileInfo) time.Time {
	stat := info.Sys().(*syscall.Stat_t)
	return time.Unix(stat.Ctim.Sec, stat.Ctim.Nsec)
}

// getLastAccessTime retrieves the last access time of a file.
// It extracts the access time from the `Atim` field of the `syscall.Stat_t`
// structure, which is obtained from the file's underlying system data.
func getLastAccessTime(info os.FileInfo) time.Time {
	stat := info.Sys().(*syscall.Stat_t)
	return time.Unix(stat.Atim.Sec, stat.Atim.Nsec)
}

// getLastWriteTime retrieves the last modification time of a file.
// It uses the `Mtim` field of the `syscall.Stat_t` structure to obtain the
// modification time.
func getLastWriteTime(info os.FileInfo) time.Time {
	stat := info.Sys().(*syscall.Stat_t)
	return time.Unix(stat.Mtim.Sec, stat.Mtim.Nsec)
}

// GetSize returns the size of a file or directory. On Unix-like systems,
// the size reported by `os.FileInfo` is generally accurate for both files
// and directories, so this function simply returns that value.
// It returns the size of the file. For directories, it returns 0 as directory
// size is not consistently defined on Unix-like systems.
// For other file types, it returns the size from the os.FileInfo.
func GetSize(info os.FileInfo, path string) int64 {
	return info.Size()
}
