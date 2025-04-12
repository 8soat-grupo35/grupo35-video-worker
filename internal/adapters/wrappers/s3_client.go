package wrappers

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

//go:generate mockgen -source=s3_client.go -destination=mock/s3_client.go
type IS3Client interface {
	Download(ctx context.Context, w io.WriterAt, input *s3.GetObjectInput, options ...func(*manager.Downloader)) (n int64, err error)
	Upload(ctx context.Context, input *s3.PutObjectInput, opts ...func(*manager.Uploader)) (*manager.UploadOutput, error)
}

type S3Client struct {
	client *s3.Client
}

// Download implements IS3Client.
func (s S3Client) Download(ctx context.Context, w io.WriterAt, input *s3.GetObjectInput, options ...func(*manager.Downloader)) (n int64, err error) {
	downloader := manager.NewDownloader(s.client)
	return downloader.Download(ctx, w, input, options...)
}

// Upload implements IS3Client.
func (s S3Client) Upload(ctx context.Context, input *s3.PutObjectInput, opts ...func(*manager.Uploader)) (*manager.UploadOutput, error) {
	uploader := manager.NewUploader(s.client)
	return uploader.Upload(ctx, input, opts...)
}

func NewS3Client(cfg aws.Config) IS3Client {
	return S3Client{
		client: s3.NewFromConfig(cfg),
	}
}
