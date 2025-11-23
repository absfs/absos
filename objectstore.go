// Package absos defines an abstract object storage interface for cloud object stores.
// It provides a unified API for interacting with object storage systems like AWS S3,
// Google Cloud Storage, Azure Blob Storage, and others.
package absos

import "context"

// ObjectStore represents a top-level object storage service.
// It provides methods for managing buckets (containers) within the storage system.
type ObjectStore interface {
	// CreateBucket creates a new bucket with the specified name.
	// Returns an error if the bucket already exists or creation fails.
	CreateBucket(ctx context.Context, bucket string) error

	// DeleteBucket deletes the bucket with the specified name.
	// Returns an error if the bucket doesn't exist, is not empty, or deletion fails.
	DeleteBucket(ctx context.Context, bucket string) error

	// ListBuckets returns a list of all buckets in the object store.
	// Returns an error if the operation fails.
	ListBuckets(ctx context.Context) ([]Bucket, error)
}
