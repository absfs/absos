package absos

import (
	"context"
	"io"
	"time"
)

// Object represents a stored object in an object store.
// It provides access to object metadata and contents.
type Object interface {
	// Bucket returns the name of the bucket containing this object.
	Bucket() string

	// Key returns the unique key (path) of the object within its bucket.
	Key() string

	// Size returns the size of the object in bytes.
	Size() int64

	// ModTime returns the last modification time of the object.
	ModTime() time.Time

	// AccessTime returns the last access time of the object.
	AccessTime() time.Time

	// ETag returns the entity tag (usually the MD5 hash) of the object.
	ETag() []byte

	// StorageClass returns the storage class of the object (e.g., STANDARD, GLACIER).
	StorageClass() string

	// Head retrieves the full metadata for this object without downloading its contents.
	Head(ctx context.Context) (ObjectHeader, error)

	// Open opens the object for reading and returns a reader for its contents.
	Open(ctx context.Context) (io.ReadCloser, error)
}

// ObjectHeader represents the complete metadata for an object.
// It includes all information about the object without its actual contents.
type ObjectHeader interface {
	// Bucket returns the name of the bucket containing this object.
	Bucket() string

	// Key returns the unique key (path) of the object within its bucket.
	Key() string

	// Size returns the size of the object in bytes.
	Size() int64

	// ModTime returns the last modification time of the object.
	ModTime() time.Time

	// AccessTime returns the last access time of the object.
	AccessTime() time.Time

	// ETag returns the entity tag (usually the MD5 hash) of the object.
	ETag() []byte

	// MimeType returns the MIME type (Content-Type) of the object.
	MimeType() string

	// Metadata returns custom user-defined metadata associated with the object.
	Metadata() map[string]string

	// Version returns the version ID of the object if versioning is enabled.
	Version() string

	// Redirect returns the URL to redirect to if this object is a redirect.
	Redirect() string

	// ServerSideEncryption returns the server-side encryption details for the object.
	ServerSideEncryption() *SSE

	// StorageClass returns the storage class of the object (e.g., STANDARD, GLACIER).
	StorageClass() string
}

// SSE represents server-side encryption configuration for an object.
type SSE struct {
	// Algorithms specifies the encryption algorithm used.
	Algorithms string `json:"algorithms"`

	// KeyMD5 is the MD5 hash of the encryption key.
	KeyMD5 string `json:"key_md5"`

	// KMSKeyId is the AWS KMS key ID used for encryption.
	KMSKeyId string `json:"kms_key_id"`

	// ServerSideEncryption specifies the server-side encryption method.
	ServerSideEncryption string `json:"server_side_encryption"`
}
