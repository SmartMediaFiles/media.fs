//go:build !windows

package fs

import (
	"os"
	"syscall"
	"time"
)

func getCreationTime(info os.FileInfo) time.Time {
	// Unix systems typically don't store creation time, return zero time
	return time.Time{}
}

func getLastAccessTime(info os.FileInfo) time.Time {
	stat := info.Sys().(*syscall.Stat_t)
	return time.Unix(stat.Atim.Sec, stat.Atim.Nsec)
}

func getLastWriteTime(info os.FileInfo) time.Time {
	stat := info.Sys().(*syscall.Stat_t)
	return time.Unix(stat.Mtim.Sec, stat.Mtim.Nsec)
}
