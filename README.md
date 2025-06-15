# 📁 media.fs

![Work in Progress](https://img.shields.io/badge/Status-Work%20in%20Progress-yellow)  
[![Go Report Card](https://goreportcard.com/badge/github.com/smartmediafiles/media.fs)](https://goreportcard.com/report/github.com/smartmediafiles/media.fs)
[![GoDoc](https://pkg.go.dev/badge/github.com/smartmediafiles/media.fs/fs)](https://pkg.go.dev/github.com/smartmediafiles/media.fs/fs)
[![Release](https://img.shields.io/github/release/smartmediafiles/media.fs.svg?style=flat)](https://github.com/smartmediafiles/media.fs/releases)


## 📝 Overview

[`media.fs`](https://github.com/SmartMediaFiles/media.fs) is a Go library designed to provide a cross-platform file system utility package. It offers a comprehensive abstraction for file metadata, allowing seamless operations across different operating systems. The library includes platform-specific implementations to handle file timestamps and other metadata efficiently.

This library is a component of the [`smartmediafiles`](https://github.com/SmartMediaFiles) organization, a suite of tools dedicated to the management and manipulation of media files, including images, videos, and more.

### ✨ Features

- **Cross-Platform Support:** Utilizes build tags to provide OS-specific implementations for file operations.
- **File Metadata Abstraction:** Implements a `FileInfo` interface that extends the standard Go `fs.FileInfo` interface with additional methods for retrieving file path, title, extension, and timestamps.
- **Utility Functions:** Includes functions to check if a file or directory exists, is empty, and resolve file paths, including symbolic links.


## 🚀 Installation

To install the package, use the following command:

```bash
go get -u github.com/smartmediafiles/media.fs/fs
```

## 💻 Usage

Here's a basic example of how to use the `media.fs` package:

```go
package main

import (
	"fmt"
	"log"

	"github.com/smartmediafiles/media.fs/fs"
)

func main() {
	// Replace with the path to your file
	info, err := fs.NewFileInfo("path/to/your/file.ext")
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

## 📜 License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

---

⚠️ **Note:** This README will be updated regularly as the project progresses. Check back often for the latest information!
