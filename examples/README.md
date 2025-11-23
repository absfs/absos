# AbsOS Examples

This directory contains example implementations of the AbsOS interfaces.

## Memory Store

The `memory` package provides a simple in-memory implementation of the AbsOS interfaces. This is useful for:

- Testing your code that uses AbsOS
- Understanding how to implement the interfaces
- Prototyping without cloud dependencies

### Usage

```go
package main

import (
    "context"
    "fmt"
    "log"
    "strings"

    "github.com/absfs/absos/examples/memory"
)

func main() {
    ctx := context.Background()

    // Create a new in-memory store
    store := memory.NewStore()

    // Create a bucket
    if err := store.CreateBucket(ctx, "my-bucket"); err != nil {
        log.Fatal(err)
    }

    // List buckets
    buckets, err := store.ListBuckets(ctx)
    if err != nil {
        log.Fatal(err)
    }

    bucket := buckets[0]
    fmt.Printf("Created bucket: %s\n", bucket.Name())

    // Upload an object
    data := strings.NewReader("Hello, AbsOS!")
    if err := bucket.Put(ctx, "greeting.txt", data); err != nil {
        log.Fatal(err)
    }

    // Download the object
    reader, err := bucket.Get(ctx, "greeting.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer reader.Close()

    // Read the content
    content := make([]byte, 13)
    if _, err := reader.Read(content); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Object content: %s\n", string(content))
}
```

## Implementing Your Own Provider

To implement support for a cloud provider:

1. Create types that implement the `absos.ObjectStore`, `absos.Bucket`, `absos.Object`, and `absos.ObjectHeader` interfaces
2. Add proper error handling using the error types from `absos` package
3. Ensure thread-safety if your implementation will be used concurrently
4. Add comprehensive tests
5. Document provider-specific behavior and requirements

See the `memory` package for a reference implementation.
