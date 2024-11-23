# media.fs

![Work in Progress](https://img.shields.io/badge/Status-Work%20in%20Progress-yellow)  
[![Go Report Card](https://goreportcard.com/badge/github.com/SmartMediaFiles/media.fs)](https://goreportcard.com/report/github.com/SmartMediaFiles/media.fs)
[![GoDoc](https://pkg.go.dev/badge/github.com/SmartMediaFiles/media.fs)](https://pkg.go.dev/github.com/SmartMediaFiles/media.fs)
[![Release](https://img.shields.io/github/release/SmartMediaFiles/media.fs.svg?style=flat)](https://github.com/SmartMediaFiles/media.fs/releases)


## Overview

`media.fs` is a Go library designed to provide a cross-platform file system utility package. It offers a comprehensive abstraction for file metadata, allowing seamless operations across different operating systems. The library includes platform-specific implementations to handle file timestamps and other metadata efficiently.

### Features

- **Cross-Platform Support:** Utilizes build tags to provide OS-specific implementations for file operations.
- **File Metadata Abstraction:** Implements a `FileInfo` interface that extends the standard Go `fs.FileInfo` interface with additional methods for retrieving file path, title, extension, and timestamps.
- **Utility Functions:** Includes functions to check if a file or directory exists, is empty, and resolve file paths, including symbolic links.


## Installation

To install the package, use the following command:

```bash
go get -u github.com/SmartMediaFiles/media.fs
```

## Usage

Here's a basic example of how to use the `media.fs` package:

```go
package main
import (
    "fmt"
    "github.com/smartmediafiles/media.fs"
)
func main() {
    info, err := fs.NewFileInfo("path/to/your/file")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
	
    fmt.Println("File Name:", info.Name())
    fmt.Println("File Path:", info.Path())
    fmt.Println("File Size:", info.Size())
    fmt.Println("Is Directory:", info.IsDir())
}
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

---

⚠️ **Note:** This README will be updated regularly as the project progresses. Check back often for the latest information!
