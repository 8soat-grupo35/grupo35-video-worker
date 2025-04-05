package gateways

import (
	"context"
	"errors"
	"grupo35-video-worker/internal/interfaces/repository"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3 struct {
	client     *s3.Client
	bucketName *string
}

func NewS3Manager(cfg aws.Config) repository.S3 {
	s3Instance := s3.NewFromConfig(cfg)

	return &S3{
		client: s3Instance,
	}
}

func (S *S3) SetBucketName(bucketName string) {
	S.bucketName = &bucketName
}

func (S *S3) DownloadFile(key string, destinationPath string) error {
	if S.bucketName == nil {
		return errors.New("bucket name is not set")
	}

	f, err := os.Create(destinationPath)

	if err != nil {
		return err
	}

	defer f.Close()

	downloader := manager.NewDownloader(S.client)

	_, err = downloader.Download(context.TODO(), f, &s3.GetObjectInput{
		Bucket: aws.String(*S.bucketName),
		Key:    aws.String(key),
	})

	return err
}

func (S *S3) UploadFile(key string, filePath string) error {
	if S.bucketName == nil {
		return errors.New("bucket name is not set")
	}

	f, err := os.Open(filePath)

	if err != nil {
		return err
	}

	defer f.Close()

	uploader := manager.NewUploader(S.client)

	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(*S.bucketName),
		Key:    aws.String(key),
		Body:   f,
	})

	return err
}
