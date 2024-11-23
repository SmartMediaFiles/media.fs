//go:build windows

package fs

import (
	"os"
	"syscall"
	"time"
)

// This file provides Windows-specific implementations for retrieving file timestamps.
// It is part of a cross-platform file system utility package, where different
// implementations are provided for different operating systems using build tags.

// getCreationTime returns the creation time of a file.
// On Windows systems, the creation time is accessible via the syscall.Win32FileAttributeData
// structure. This function extracts the creation time from the CreationTime field
// and returns it as a time.Time object.
func getCreationTime(info os.FileInfo) time.Time {
	stat := info.Sys().(*syscall.Win32FileAttributeData)
	return time.Unix(0, stat.CreationTime.Nanoseconds())
}

// getLastAccessTime retrieves the last access time of a file.
// It uses the syscall.Win32FileAttributeData structure to obtain the access time.
// The LastAccessTime field provides the necessary nanoseconds, which are used to
// construct a time.Time object representing the last access time.
func getLastAccessTime(info os.FileInfo) time.Time {
	stat := info.Sys().(*syscall.Win32FileAttributeData)
	return time.Unix(0, stat.LastAccessTime.Nanoseconds())
}

// getLastWriteTime retrieves the last modification time of a file.
// Similar to getLastAccessTime, it uses the syscall.Win32FileAttributeData structure
// to obtain the modification time. The LastWriteTime field provides the necessary
// nanoseconds, which are used to construct a time.Time object representing the last
// modification time.
func getLastWriteTime(info os.FileInfo) time.Time {
	stat := info.Sys().(*syscall.Win32FileAttributeData)
	return time.Unix(0, stat.LastWriteTime.Nanoseconds())
}
