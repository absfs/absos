package absos

import (
	"errors"
	"testing"
)

func TestBucketError(t *testing.T) {
	bucket := "test-bucket"
	baseErr := ErrBucketNotFound

	err := &BucketError{
		Bucket: bucket,
		Err:    baseErr,
	}

	// Test Error() method
	expected := `bucket "test-bucket": bucket not found`
	if err.Error() != expected {
		t.Errorf("expected %q, got %q", expected, err.Error())
	}

	// Test Unwrap() method
	if !errors.Is(err, baseErr) {
		t.Error("errors.Is should return true for base error")
	}

	unwrapped := err.Unwrap()
	if unwrapped != baseErr {
		t.Errorf("expected unwrapped error to be %v, got %v", baseErr, unwrapped)
	}
}

func TestObjectError(t *testing.T) {
	bucket := "test-bucket"
	key := "test-key"
	baseErr := ErrObjectNotFound

	err := &ObjectError{
		Bucket: bucket,
		Key:    key,
		Err:    baseErr,
	}

	// Test Error() method
	expected := `object "test-key" in bucket "test-bucket": object not found`
	if err.Error() != expected {
		t.Errorf("expected %q, got %q", expected, err.Error())
	}

	// Test Unwrap() method
	if !errors.Is(err, baseErr) {
		t.Error("errors.Is should return true for base error")
	}

	unwrapped := err.Unwrap()
	if unwrapped != baseErr {
		t.Errorf("expected unwrapped error to be %v, got %v", baseErr, unwrapped)
	}
}

func TestErrorConstants(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected string
	}{
		{"BucketNotFound", ErrBucketNotFound, "bucket not found"},
		{"BucketAlreadyExists", ErrBucketAlreadyExists, "bucket already exists"},
		{"BucketNotEmpty", ErrBucketNotEmpty, "bucket not empty"},
		{"ObjectNotFound", ErrObjectNotFound, "object not found"},
		{"InvalidKey", ErrInvalidKey, "invalid object key"},
		{"PermissionDenied", ErrPermissionDenied, "permission denied"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err.Error() != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, tt.err.Error())
			}
		})
	}
}
