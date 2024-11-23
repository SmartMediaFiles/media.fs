//go:build !windows

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
// On Unix systems, the creation time is typically not stored or accessible.
// Therefore, this function returns a zero time value, indicating that the
// creation time is not available. This is a limitation of Unix file systems
// compared to Windows, where creation time is usually available.
func getCreationTime(info os.FileInfo) time.Time {
	// Unix systems typically don't store creation time, return zero time
	return time.Time{}
}

// getLastAccessTime retrieves the last access time of a file.
// It extracts the access time from the syscall.Stat_t structure, which is
// obtained from the file's underlying system data. The access time is returned
// as a time.Time object, constructed from the seconds and nanoseconds fields
// of the Atim member of Stat_t.
func getLastAccessTime(info os.FileInfo) time.Time {
	stat := info.Sys().(*syscall.Stat_t)
	return time.Unix(stat.Atim.Sec, stat.Atim.Nsec)
}

// getLastWriteTime retrieves the last modification time of a file.
// Similar to getLastAccessTime, it uses the syscall.Stat_t structure to
// obtain the modification time. The Mtim member of Stat_t provides the
// necessary seconds and nanoseconds, which are used to construct a time.Time
// object representing the last modification time.
func getLastWriteTime(info os.FileInfo) time.Time {
	stat := info.Sys().(*syscall.Stat_t)
	return time.Unix(stat.Mtim.Sec, stat.Mtim.Nsec)
}
