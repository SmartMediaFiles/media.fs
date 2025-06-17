# 📁 media.fs

[![Work in Progress](https://img.shields.io/badge/Status-Work%20in%20Progress-yellow)](https://shields.io)
[![Go Report Card](https://goreportcard.com/badge/github.com/smartmediafiles/media.fs)](https://goreportcard.com/report/github.com/smartmediafiles/media.fs)
[![GoDoc](https://pkg.go.dev/badge/github.com/smartmediafiles/media.fs/fs)](https://pkg.go.dev/github.com/smartmediafiles/media.fs/fs)
[![Release](https://img.shields.io/github/release/smartmediafiles/media.fs.svg?style=flat)](https://github.com/smartmediafiles/media.fs/releases)

## Overview

`media.fs` is a low-level utility library for the **SmartMediaFiles ecosystem**. It provides a set of functions to interact with the file system in a reliable and cross-platform manner. It offers a comprehensive abstraction for file metadata, allowing seamless operations across different operating systems.

## Features

- **Cross-Platform Support:** Utilizes build tags to provide OS-specific implementations for file operations.
- **File Metadata Abstraction:** Implements a `FileInfo` interface that extends the standard Go `fs.FileInfo`.
- **Path Verification:** Includes functions to check if a path points to a file (`IsFile`) or a directory (`IsDir`).
- **Content-Check:** Includes a function to check if a file or directory is empty (`IsEmpty`).
- **Path Resolution:** Includes a function to resolve a path to an absolute path, expanding tildes `~` and evaluating symbolic links (`Resolve`).

## Installation

```bash
go get -u github.com/smartmediafiles/media.fs/fs
```

## Usage

Here's a basic example of how to use the `media.fs` package:

```go
package main

import (
	"fmt"
	"log"
	"github.com/smartmediafiles/media.fs/fs"
)

func main() {
	// Resolve a path
	resolvedPath, err := fs.Resolve("~/your-file.txt")
	if err != nil {
		// On error, it's likely the file doesn't exist
		log.Fatal(err)
	}
	fmt.Printf("Absolute Path: %s\n", resolvedPath)

	// Get file information
	info, err := fs.NewFileInfo(resolvedPath)
	if err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Println("Name:", info.Name())
	fmt.Println("Title:", info.Title())
	fmt.Println("Extension:", info.Ext())
	fmt.Println("Path:", info.Path())
	fmt.Println("Absolute Path:", info.Abs())
	fmt.Println("Size (bytes):", info.Size())
	fmt.Println("Is Directory:", info.IsDir())
	fmt.Println("Creation Time:", info.CreationTime())
	fmt.Println("Last Access Time:", info.LastAccessTime())
	fmt.Println("Last Write Time:", info.LastWriteTime())
}
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

---

⚠️ **Note:** This README will be updated regularly as the project progresses. Check back often for the latest information!
