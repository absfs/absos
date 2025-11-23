package absos

import (
	"context"
	"io"
	"time"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Owner represents the owner of a bucket or object.
type Owner interface {
	// Name returns the display name of the owner.
	Name() string

	// ID returns the unique identifier of the owner.
	ID() string
}

// Page represents a paginated list of objects and prefixes within a bucket.
// It supports iterating through large result sets.
type Page interface {
	// Objects returns the list of objects in this page.
	Objects() []Object

	// Prefixes returns the list of common prefixes (virtual directories) in this page.
	Prefixes() []string

	// NextPage returns the token to use for fetching the next page.
	// Returns an empty string if this is the last page.
	NextPage() string

	// Last returns true if this is the last page of results.
	Last() bool
}

// Bucket represents a storage bucket (container) within an object store.
// It provides methods for managing objects within the bucket.
type Bucket interface {
	// Name returns the name of the bucket.
	Name() string

	// CreationTime returns the time when the bucket was created.
	CreationTime() time.Time

	// Owner returns the owner of the bucket.
	Owner() Owner

	// ObjectPage returns a paginated list of objects with the given prefix and delimiter.
	// The token parameter is used for pagination; pass an empty string for the first page.
	ObjectPage(ctx context.Context, prefix, delimiter, token string) (Page, error)

	// Head retrieves metadata for the object with the specified key without downloading the object.
	Head(ctx context.Context, key string) (ObjectHeader, error)

	// PutBatch uploads multiple objects in batch using the provided iterator.
	PutBatch(ctx context.Context, iter s3manager.BatchUploadIterator) error

	// Put uploads an object with the specified key from the provided reader.
	Put(ctx context.Context, key string, data io.ReadSeeker) error

	// Get retrieves the object with the specified key and returns a reader for its contents.
	Get(ctx context.Context, key string) (io.ReadCloser, error)

	// Delete removes the object with the specified key from the bucket.
	Delete(ctx context.Context, key string) error
}
