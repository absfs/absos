// Package memory provides a simple in-memory implementation of the absos interfaces.
// This is intended for testing and demonstration purposes only.
package memory

import (
	"bytes"
	"context"
	"io"
	"sync"
	"time"

	"github.com/absfs/absos"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Store is an in-memory implementation of absos.ObjectStore.
type Store struct {
	mu      sync.RWMutex
	buckets map[string]*Bucket
}

// NewStore creates a new in-memory object store.
func NewStore() *Store {
	return &Store{
		buckets: make(map[string]*Bucket),
	}
}

// CreateBucket creates a new bucket in memory.
func (s *Store) CreateBucket(ctx context.Context, name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.buckets[name]; exists {
		return &absos.BucketError{Bucket: name, Err: absos.ErrBucketAlreadyExists}
	}

	s.buckets[name] = &Bucket{
		name:    name,
		created: time.Now(),
		objects: make(map[string]*object),
	}

	return nil
}

// DeleteBucket deletes a bucket from memory.
func (s *Store) DeleteBucket(ctx context.Context, name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	bucket, exists := s.buckets[name]
	if !exists {
		return &absos.BucketError{Bucket: name, Err: absos.ErrBucketNotFound}
	}

	bucket.mu.RLock()
	notEmpty := len(bucket.objects) > 0
	bucket.mu.RUnlock()

	if notEmpty {
		return &absos.BucketError{Bucket: name, Err: absos.ErrBucketNotEmpty}
	}

	delete(s.buckets, name)
	return nil
}

// ListBuckets returns all buckets in the store.
func (s *Store) ListBuckets(ctx context.Context) ([]absos.Bucket, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	buckets := make([]absos.Bucket, 0, len(s.buckets))
	for _, b := range s.buckets {
		buckets = append(buckets, b)
	}

	return buckets, nil
}

// Bucket is an in-memory implementation of absos.Bucket.
type Bucket struct {
	mu      sync.RWMutex
	name    string
	created time.Time
	objects map[string]*object
}

// Name returns the bucket name.
func (b *Bucket) Name() string {
	return b.name
}

// CreationTime returns when the bucket was created.
func (b *Bucket) CreationTime() time.Time {
	return b.created
}

// Owner returns nil for this implementation.
func (b *Bucket) Owner() absos.Owner {
	return nil
}

// ObjectPage returns a page of objects (simplified pagination).
func (b *Bucket) ObjectPage(ctx context.Context, prefix, delimiter, token string) (absos.Page, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	// Simplified implementation - returns all objects matching prefix
	var objects []absos.Object
	for key, obj := range b.objects {
		if prefix == "" || len(key) >= len(prefix) && key[:len(prefix)] == prefix {
			objects = append(objects, obj)
		}
	}

	return &page{objects: objects, last: true}, nil
}

// Head retrieves object metadata.
func (b *Bucket) Head(ctx context.Context, key string) (absos.ObjectHeader, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	obj, exists := b.objects[key]
	if !exists {
		return nil, &absos.ObjectError{Bucket: b.name, Key: key, Err: absos.ErrObjectNotFound}
	}

	return obj, nil
}

// PutBatch is not implemented in this simple example.
func (b *Bucket) PutBatch(ctx context.Context, iter s3manager.BatchUploadIterator) error {
	return nil
}

// Put stores an object in memory.
func (b *Bucket) Put(ctx context.Context, key string, data io.ReadSeeker) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	content, err := io.ReadAll(data)
	if err != nil {
		return err
	}

	b.objects[key] = &object{
		bucket:  b.name,
		key:     key,
		data:    content,
		modTime: time.Now(),
	}

	return nil
}

// Get retrieves an object from memory.
func (b *Bucket) Get(ctx context.Context, key string) (io.ReadCloser, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	obj, exists := b.objects[key]
	if !exists {
		return nil, &absos.ObjectError{Bucket: b.name, Key: key, Err: absos.ErrObjectNotFound}
	}

	return io.NopCloser(bytes.NewReader(obj.data)), nil
}

// Delete removes an object from memory.
func (b *Bucket) Delete(ctx context.Context, key string) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, exists := b.objects[key]; !exists {
		return &absos.ObjectError{Bucket: b.name, Key: key, Err: absos.ErrObjectNotFound}
	}

	delete(b.objects, key)
	return nil
}

type object struct {
	bucket  string
	key     string
	data    []byte
	modTime time.Time
}

func (o *object) Bucket() string                   { return o.bucket }
func (o *object) Key() string                      { return o.key }
func (o *object) Size() int64                      { return int64(len(o.data)) }
func (o *object) ModTime() time.Time               { return o.modTime }
func (o *object) AccessTime() time.Time            { return o.modTime }
func (o *object) ETag() []byte                     { return nil }
func (o *object) StorageClass() string             { return "STANDARD" }
func (o *object) MimeType() string                 { return "application/octet-stream" }
func (o *object) Metadata() map[string]string      { return nil }
func (o *object) Version() string                  { return "" }
func (o *object) Redirect() string                 { return "" }
func (o *object) ServerSideEncryption() *absos.SSE { return nil }

func (o *object) Head(ctx context.Context) (absos.ObjectHeader, error) {
	return o, nil
}

func (o *object) Open(ctx context.Context) (io.ReadCloser, error) {
	return io.NopCloser(bytes.NewReader(o.data)), nil
}

type page struct {
	objects []absos.Object
	last    bool
}

func (p *page) Objects() []absos.Object { return p.objects }
func (p *page) Prefixes() []string      { return nil }
func (p *page) NextPage() string        { return "" }
func (p *page) Last() bool              { return p.last }
