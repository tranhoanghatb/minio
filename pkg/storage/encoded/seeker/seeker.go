package seeker

import (
	"io"

	"github.com/minio-io/minio/pkg/encoding/erasure"
	"github.com/minio-io/minio/pkg/storage"
	donuterasure "github.com/minio-io/minio/pkg/storage/donut/erasure"
)

type Seeker interface {
	ListBuckets() ([]storage.BucketMetadata, error)

	GetReader(bucket, object string, chunk uint, part uint8) (donuterasure.DataHeader, io.Reader, error)
	// TODO this should probably write async and return via a channel. For now it blocks.
	Write(bucket, object string, chunk int, part uint8, length int, params erasure.EncoderParams, reader io.Reader) error

	GetObjectMetadata(bucket string, object string, prefix string) (storage.ObjectMetadata, error)
	ListObjects(bucket string, resources storage.BucketResourcesMetadata) ([]storage.ObjectMetadata, storage.BucketResourcesMetadata, error)

	SetPolicy(bucket string, policy storage.BucketPolicy) error
	GetPolicy(bucket string) (storage.BucketPolicy, error)

	CreateBucket(bucket string) error
}
