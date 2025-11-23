package memory

import (
	"context"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/absfs/absos"
)

func TestStoreCreateBucket(t *testing.T) {
	store := NewStore()
	ctx := context.Background()

	// Test creating a new bucket
	err := store.CreateBucket(ctx, "test-bucket")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Test creating duplicate bucket
	err = store.CreateBucket(ctx, "test-bucket")
	if err == nil {
		t.Fatal("expected error for duplicate bucket")
	}

	if !errors.Is(err, absos.ErrBucketAlreadyExists) {
		t.Errorf("expected ErrBucketAlreadyExists, got %v", err)
	}
}

func TestStoreListBuckets(t *testing.T) {
	store := NewStore()
	ctx := context.Background()

	// Create multiple buckets
	bucketNames := []string{"bucket1", "bucket2", "bucket3"}
	for _, name := range bucketNames {
		if err := store.CreateBucket(ctx, name); err != nil {
			t.Fatalf("failed to create bucket %s: %v", name, err)
		}
	}

	// List buckets
	buckets, err := store.ListBuckets(ctx)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(buckets) != len(bucketNames) {
		t.Errorf("expected %d buckets, got %d", len(bucketNames), len(buckets))
	}
}

func TestStoreDeleteBucket(t *testing.T) {
	store := NewStore()
	ctx := context.Background()

	// Test deleting non-existent bucket
	err := store.DeleteBucket(ctx, "non-existent")
	if err == nil {
		t.Fatal("expected error for non-existent bucket")
	}

	if !errors.Is(err, absos.ErrBucketNotFound) {
		t.Errorf("expected ErrBucketNotFound, got %v", err)
	}

	// Create a bucket
	if err := store.CreateBucket(ctx, "test-bucket"); err != nil {
		t.Fatalf("failed to create bucket: %v", err)
	}

	// Delete empty bucket
	err = store.DeleteBucket(ctx, "test-bucket")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestStoreDeleteNonEmptyBucket(t *testing.T) {
	store := NewStore()
	ctx := context.Background()

	// Create bucket and add object
	if err := store.CreateBucket(ctx, "test-bucket"); err != nil {
		t.Fatalf("failed to create bucket: %v", err)
	}

	buckets, _ := store.ListBuckets(ctx)
	bucket := buckets[0]

	data := strings.NewReader("test data")
	if err := bucket.Put(ctx, "test-key", data); err != nil {
		t.Fatalf("failed to put object: %v", err)
	}

	// Try to delete non-empty bucket
	err := store.DeleteBucket(ctx, "test-bucket")
	if err == nil {
		t.Fatal("expected error for non-empty bucket")
	}

	if !errors.Is(err, absos.ErrBucketNotEmpty) {
		t.Errorf("expected ErrBucketNotEmpty, got %v", err)
	}
}

func TestBucketPutGet(t *testing.T) {
	store := NewStore()
	ctx := context.Background()

	// Create bucket
	if err := store.CreateBucket(ctx, "test-bucket"); err != nil {
		t.Fatalf("failed to create bucket: %v", err)
	}

	buckets, _ := store.ListBuckets(ctx)
	bucket := buckets[0]

	// Put object
	testData := "Hello, World!"
	data := strings.NewReader(testData)
	if err := bucket.Put(ctx, "test-key", data); err != nil {
		t.Fatalf("failed to put object: %v", err)
	}

	// Get object
	reader, err := bucket.Get(ctx, "test-key")
	if err != nil {
		t.Fatalf("failed to get object: %v", err)
	}
	defer reader.Close()

	content, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("failed to read object: %v", err)
	}

	if string(content) != testData {
		t.Errorf("expected %q, got %q", testData, string(content))
	}
}

func TestBucketHead(t *testing.T) {
	store := NewStore()
	ctx := context.Background()

	// Create bucket
	if err := store.CreateBucket(ctx, "test-bucket"); err != nil {
		t.Fatalf("failed to create bucket: %v", err)
	}

	buckets, _ := store.ListBuckets(ctx)
	bucket := buckets[0]

	// Put object
	testData := "Hello, World!"
	data := strings.NewReader(testData)
	if err := bucket.Put(ctx, "test-key", data); err != nil {
		t.Fatalf("failed to put object: %v", err)
	}

	// Head object
	header, err := bucket.Head(ctx, "test-key")
	if err != nil {
		t.Fatalf("failed to head object: %v", err)
	}

	if header.Key() != "test-key" {
		t.Errorf("expected key %q, got %q", "test-key", header.Key())
	}

	if header.Size() != int64(len(testData)) {
		t.Errorf("expected size %d, got %d", len(testData), header.Size())
	}
}

func TestBucketDelete(t *testing.T) {
	store := NewStore()
	ctx := context.Background()

	// Create bucket
	if err := store.CreateBucket(ctx, "test-bucket"); err != nil {
		t.Fatalf("failed to create bucket: %v", err)
	}

	buckets, _ := store.ListBuckets(ctx)
	bucket := buckets[0]

	// Delete non-existent object
	err := bucket.Delete(ctx, "non-existent")
	if err == nil {
		t.Fatal("expected error for non-existent object")
	}

	if !errors.Is(err, absos.ErrObjectNotFound) {
		t.Errorf("expected ErrObjectNotFound, got %v", err)
	}

	// Put and delete object
	data := strings.NewReader("test data")
	if err := bucket.Put(ctx, "test-key", data); err != nil {
		t.Fatalf("failed to put object: %v", err)
	}

	err = bucket.Delete(ctx, "test-key")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Verify object is deleted
	_, err = bucket.Get(ctx, "test-key")
	if err == nil {
		t.Fatal("expected error for deleted object")
	}
}

func TestBucketObjectPage(t *testing.T) {
	store := NewStore()
	ctx := context.Background()

	// Create bucket
	if err := store.CreateBucket(ctx, "test-bucket"); err != nil {
		t.Fatalf("failed to create bucket: %v", err)
	}

	buckets, _ := store.ListBuckets(ctx)
	bucket := buckets[0]

	// Add multiple objects
	objects := []string{"file1.txt", "file2.txt", "prefix/file3.txt"}
	for _, key := range objects {
		data := strings.NewReader("test data")
		if err := bucket.Put(ctx, key, data); err != nil {
			t.Fatalf("failed to put object %s: %v", key, err)
		}
	}

	// List all objects
	page, err := bucket.ObjectPage(ctx, "", "", "")
	if err != nil {
		t.Fatalf("failed to list objects: %v", err)
	}

	if len(page.Objects()) != len(objects) {
		t.Errorf("expected %d objects, got %d", len(objects), len(page.Objects()))
	}

	// List objects with prefix
	page, err = bucket.ObjectPage(ctx, "prefix/", "", "")
	if err != nil {
		t.Fatalf("failed to list objects: %v", err)
	}

	if len(page.Objects()) != 1 {
		t.Errorf("expected 1 object with prefix, got %d", len(page.Objects()))
	}
}
