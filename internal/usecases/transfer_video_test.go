package usecases

import (
	mock_repository "grupo35-video-worker/internal/interfaces/repository/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetVideo_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	videoPath := "video.mp4"
	videoDownloadOutputPath := "video.mp4"

	s3repo := mock_repository.NewMockS3(ctrl)
	s3repo.EXPECT().SetBucketName("grupo35-video-uploaded").Return().AnyTimes()
	s3repo.EXPECT().DownloadFile(videoPath, videoDownloadOutputPath).Return(nil).AnyTimes()

	transfer := NewTransferFile(s3repo)

	outputPath, err := transfer.GetVideo(videoPath, videoDownloadOutputPath)

	assert.NoError(t, err)
	assert.Equal(t, videoDownloadOutputPath, outputPath)
}

func TestGetVideo_Error(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	videoPath := "video.mp4"
	videoDownloadOutputPath := "video.mp4"

	s3repo := mock_repository.NewMockS3(ctrl)
	s3repo.EXPECT().SetBucketName("grupo35-video-uploaded").Return().AnyTimes()
	s3repo.EXPECT().DownloadFile(videoPath, videoDownloadOutputPath).Return(assert.AnError).AnyTimes()

	transfer := NewTransferFile(s3repo)

	outputPath, err := transfer.GetVideo(videoPath, videoDownloadOutputPath)

	assert.EqualError(t, err, assert.AnError.Error())
	assert.Equal(t, videoDownloadOutputPath, outputPath)
}

func TestUploadZip_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	zipPath := "screenshots.zip"

	s3repo := mock_repository.NewMockS3(ctrl)
	s3repo.EXPECT().SetBucketName("grupo35-video-processed").Return().AnyTimes()
	s3repo.EXPECT().UploadFile(zipPath, zipPath).Return(nil).AnyTimes()

	transfer := NewTransferFile(s3repo)

	err := transfer.UploadZip(zipPath)

	assert.NoError(t, err)
}

func TestUploadZip_error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	zipPath := "screenshots.zip"

	s3repo := mock_repository.NewMockS3(ctrl)
	s3repo.EXPECT().SetBucketName("grupo35-video-processed").Return().AnyTimes()
	s3repo.EXPECT().UploadFile(zipPath, zipPath).Return(assert.AnError).AnyTimes()
	transfer := NewTransferFile(s3repo)
	err := transfer.UploadZip(zipPath)
	assert.EqualError(t, err, assert.AnError.Error())
}
