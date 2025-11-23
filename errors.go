package absos

import (
	"errors"
	"fmt"
)

// Common errors returned by object store implementations.
var (
	// ErrBucketNotFound is returned when a bucket does not exist.
	ErrBucketNotFound = errors.New("bucket not found")

	// ErrBucketAlreadyExists is returned when attempting to create a bucket that already exists.
	ErrBucketAlreadyExists = errors.New("bucket already exists")

	// ErrBucketNotEmpty is returned when attempting to delete a non-empty bucket.
	ErrBucketNotEmpty = errors.New("bucket not empty")

	// ErrObjectNotFound is returned when an object does not exist.
	ErrObjectNotFound = errors.New("object not found")

	// ErrInvalidKey is returned when an object key is invalid.
	ErrInvalidKey = errors.New("invalid object key")

	// ErrPermissionDenied is returned when access to a resource is denied.
	ErrPermissionDenied = errors.New("permission denied")
)

// BucketError wraps an error with the bucket name for context.
type BucketError struct {
	Bucket string
	Err    error
}

// Error implements the error interface.
func (e *BucketError) Error() string {
	return fmt.Sprintf("bucket %q: %v", e.Bucket, e.Err)
}

// Unwrap returns the underlying error.
func (e *BucketError) Unwrap() error {
	return e.Err
}

// ObjectError wraps an error with bucket and object key for context.
type ObjectError struct {
	Bucket string
	Key    string
	Err    error
}

// Error implements the error interface.
func (e *ObjectError) Error() string {
	return fmt.Sprintf("object %q in bucket %q: %v", e.Key, e.Bucket, e.Err)
}

// Unwrap returns the underlying error.
func (e *ObjectError) Unwrap() error {
	return e.Err
}
