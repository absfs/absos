package absos

import (
	"io"
	"time"
)

type Object interface {
	Bucket() string
	Key() string
	Size() int64
	ModTime() time.Time
	AccessTime() time.Time
	ETag() []byte // usually the md5
	StorageClass() string

	Head() (ObjectHeader, error)
	Open() (io.ReadCloser, error)
}

type ObjectHeader interface {
	Bucket() string
	Key() string
	Size() int64
	ModTime() time.Time
	AccessTime() time.Time
	ETag() []byte
	MimeType() string
	Metadata() map[string]string
	Version() string
	Redirect() string
	ServerSideEncryption() *SSE
	StorageClass() string
}

type SSE struct {
	Algorithms           string `json:"algorithms"`
	KeyMD5               string `json:"key_md5"`
	KMSKeyId             string `json:"kms_key_id"`
	ServerSideEncryption string `json:"server_side_encryption"`
}
