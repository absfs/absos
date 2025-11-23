# AbsOS - Abstract Object Store

[![Go Reference](https://pkg.go.dev/badge/github.com/absfs/absos.svg)](https://pkg.go.dev/github.com/absfs/absos)
[![Go Report Card](https://goreportcard.com/badge/github.com/absfs/absos)](https://goreportcard.com/report/github.com/absfs/absos)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

The `absos` package defines an abstract object storage interface that provides a unified API for interacting with cloud object storage systems like AWS S3, Google Cloud Storage, Azure Blob Storage, and others.

## Features

- **Unified Interface**: Common API across different object storage providers
- **Bucket Management**: Create, delete, and list buckets
- **Object Operations**: Upload, download, delete, and list objects
- **Metadata Support**: Rich metadata and header information
- **Pagination**: Efficient handling of large object lists
- **Batch Operations**: Support for batch uploads
- **Server-Side Encryption**: Built-in encryption configuration support

## Installation

```bash
go get github.com/absfs/absos
```

## Usage

### Basic Example

```go
package main

import (
    "context"
    "fmt"
    "log"
    "strings"

    "github.com/absfs/absos"
)

func example(store absos.ObjectStore) error {
    ctx := context.Background()

    // Create a new bucket
    if err := store.CreateBucket(ctx, "my-bucket"); err != nil {
        return fmt.Errorf("create bucket: %w", err)
    }

    // List all buckets
    buckets, err := store.ListBuckets(ctx)
    if err != nil {
        return fmt.Errorf("list buckets: %w", err)
    }

    for _, bucket := range buckets {
        fmt.Printf("Bucket: %s, Created: %s\n",
            bucket.Name(), bucket.CreationTime())
    }

    return nil
}
```

### Working with Objects

```go
func uploadExample(bucket absos.Bucket) error {
    ctx := context.Background()

    // Upload an object
    data := strings.NewReader("Hello, World!")
    if err := bucket.Put(ctx, "hello.txt", data); err != nil {
        return fmt.Errorf("upload object: %w", err)
    }

    // Get object metadata
    header, err := bucket.Head(ctx, "hello.txt")
    if err != nil {
        return fmt.Errorf("get metadata: %w", err)
    }

    fmt.Printf("Object: %s, Size: %d bytes, Type: %s\n",
        header.Key(), header.Size(), header.MimeType())

    // Download an object
    reader, err := bucket.Get(ctx, "hello.txt")
    if err != nil {
        return fmt.Errorf("download object: %w", err)
    }
    defer reader.Close()

    // Use reader...

    return nil
}
```

### Listing Objects with Pagination

```go
func listExample(bucket absos.Bucket) error {
    ctx := context.Background()
    token := ""

    for {
        page, err := bucket.ObjectPage(ctx, "prefix/", "/", token)
        if err != nil {
            return fmt.Errorf("list objects: %w", err)
        }

        // Process objects
        for _, obj := range page.Objects() {
            fmt.Printf("Object: %s, Size: %d\n", obj.Key(), obj.Size())
        }

        // Process prefixes (virtual directories)
        for _, prefix := range page.Prefixes() {
            fmt.Printf("Prefix: %s\n", prefix)
        }

        if page.Last() {
            break
        }

        token = page.NextPage()
    }

    return nil
}
```

## Architecture

The package defines several key interfaces:

- **ObjectStore**: Top-level interface for managing buckets
- **Bucket**: Interface for bucket-level operations and object management
- **Object**: Represents an object with basic metadata
- **ObjectHeader**: Extended metadata for an object
- **Page**: Pagination support for listing large object sets
- **Owner**: Bucket and object ownership information
- **SSE**: Server-side encryption configuration

## Implementing a Provider

To implement support for a new object storage provider, create types that implement the `absos.ObjectStore`, `absos.Bucket`, `absos.Object`, and related interfaces.

See the [examples](examples/) directory for reference implementations.

## Contributing

We strongly encourage contributions! Please fork and submit Pull Requests, and publish any implementations you create.

New provider implementations do not need to be added to this repo, but we'd be happy to link to yours. Please open an issue or submit a Pull Request with an updated README.

For more details, see [CONTRIBUTING.md](CONTRIBUTING.md).

## License

This project is governed by the MIT License. See [LICENSE](https://github.com/absfs/absos/blob/master/LICENSE)

## Related Projects

- [absfs](https://github.com/absfs) - Abstract FileSystem interfaces

