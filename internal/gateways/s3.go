package gateways

import (
	"context"
	"errors"
	"grupo35-video-worker/internal/adapters/wrappers"
	"grupo35-video-worker/internal/interfaces/repository"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3 struct {
	client     wrappers.IS3Client
	bucketName *string
}

func NewS3Manager(client wrappers.IS3Client) repository.S3 {

	return &S3{
		client: client,
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

	_, err = S.client.Download(context.TODO(), f, &s3.GetObjectInput{
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

	_, err = S.client.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(*S.bucketName),
		Key:    aws.String(key),
		Body:   f,
	})

	return err
}
