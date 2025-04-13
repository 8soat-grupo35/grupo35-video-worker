package gateways

import (
	"context"
	mock_wrappers "grupo35-video-worker/internal/adapters/wrappers/mock"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSetBucketName(t *testing.T) {
	s3manager := S3{
		client:     nil,
		bucketName: nil,
	}

	bucketName := "bucket-name"

	s3manager.SetBucketName(bucketName)

	assert.Equal(t, bucketName, *s3manager.bucketName)
}

func TestDownloadFile_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	keyBucket := "test/key"
	destinationPath := "destination.txt"
	defer os.Remove(destinationPath)

	s3client := mock_wrappers.NewMockIS3Client(ctrl)
	s3client.EXPECT().Download(context.TODO(), gomock.Any(), &s3.GetObjectInput{
		Bucket: aws.String("bucket-name"),
		Key:    aws.String(keyBucket),
	}).Return(int64(1), nil).AnyTimes()

	s3manager := NewS3Manager(s3client)
	s3manager.SetBucketName("bucket-name")

	err := s3manager.DownloadFile(keyBucket, destinationPath)

	assert.NoError(t, err)
}

func TestDownloadFile_Error_Bucket(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	keyBucket := "test/key"
	destinationPath := "destination.txt"
	//defer os.Remove(destinationPath)

	s3client := mock_wrappers.NewMockIS3Client(ctrl)

	s3manager := NewS3Manager(s3client)
	err := s3manager.DownloadFile(keyBucket, destinationPath)

	assert.EqualError(t, err, "bucket name is not set")
}

func TestDownloadFile_Error_Download(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	keyBucket := "test/key"
	destinationPath := "destination.txt"
	defer os.Remove(destinationPath)

	s3client := mock_wrappers.NewMockIS3Client(ctrl)
	s3client.EXPECT().Download(context.TODO(), gomock.Any(), &s3.GetObjectInput{
		Bucket: aws.String("bucket-name"),
		Key:    aws.String(keyBucket),
	}).Return(int64(1), assert.AnError).AnyTimes()

	s3manager := NewS3Manager(s3client)
	s3manager.SetBucketName("bucket-name")
	err := s3manager.DownloadFile(keyBucket, destinationPath)

	assert.EqualError(t, err, assert.AnError.Error())
}

func TestUploadFile_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	keyBucket := "test/key"
	destinationPath := "destination.txt"
	f, _ := os.Create(destinationPath)
	f.Close()
	defer os.Remove(destinationPath)

	s3client := mock_wrappers.NewMockIS3Client(ctrl)
	s3client.EXPECT().Upload(context.TODO(), gomock.Any()).Return(nil, nil).AnyTimes()

	s3manager := NewS3Manager(s3client)
	s3manager.SetBucketName("bucket-name")
	err := s3manager.UploadFile(keyBucket, destinationPath)
	assert.NoError(t, err)
}

func TestUploadFile_Error_Bucket(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	keyBucket := "test/key"
	destinationPath := "destination.txt"

	s3client := mock_wrappers.NewMockIS3Client(ctrl)

	s3manager := NewS3Manager(s3client)
	err := s3manager.UploadFile(keyBucket, destinationPath)

	assert.EqualError(t, err, "bucket name is not set")
}

func TestUploadFile_Error_Upload(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	keyBucket := "test/key"
	destinationPath := "destination.txt"
	f, _ := os.Create(destinationPath)
	f.Close()
	defer os.Remove(destinationPath)

	s3client := mock_wrappers.NewMockIS3Client(ctrl)
	s3client.EXPECT().Upload(context.TODO(), gomock.Any()).Return(nil, assert.AnError).AnyTimes()

	s3manager := NewS3Manager(s3client)
	s3manager.SetBucketName("bucket-name")
	err := s3manager.UploadFile(keyBucket, destinationPath)
	assert.EqualError(t, err, assert.AnError.Error())
}
