package absos

import (
	"io"
	"time"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Owner interface {
	Name() string
	ID() string
}

type Page interface {
	Objects() []Object
	Prefixes() []string
	NextPage() string
	Last() bool
}

type Bucket interface {
	Name() string
	CreationTime() time.Time
	Owner() Owner
	ObjectPage(prefix, delimiter, token string) (Page, error)
	Head(string) (ObjectHeader, error)

	PutBatch(iter s3manager.BatchUploadIterator) error
	Put(string, io.ReadSeeker) error
	Get(string) (io.ReadCloser, error)
	Delete(string) error
}
