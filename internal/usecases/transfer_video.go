package usecases

import (
	"fmt"
	"grupo35-video-worker/internal/interfaces/repository"
)

//go:generate mockgen -source=transfer_video.go -destination=mock/transfer_video.go
type ITransferFile interface {
	GetVideo(videoPath string, videoDownloadOutputPath string) (outputPath string, err error)
	UploadZip(zipPath string) (err error)
}

type TransferFile struct {
	S3  repository.S3
}

func NewTransferFile(s3 repository.S3) ITransferFile {
	return TransferFile{
		S3: s3,
	}
}

func (t TransferFile) GetVideo(videoPath string, videoDownloadOutputPath string) (outputPath string, err error) {
	fmt.Println("Getting file from S3: ", videoPath)

	t.S3.SetBucketName("grupo35-video-uploaded")

	outputPath = videoDownloadOutputPath
	err = t.S3.DownloadFile(videoPath, outputPath)

	if err != nil {
		fmt.Println(err)
	}

	return
}

func (t TransferFile) UploadZip(zipPath string) (err error) {
	fmt.Println("Uploading processed video to S3")
	t.S3.SetBucketName("grupo35-video-processed")
	err = t.S3.UploadFile(zipPath, zipPath)

	if err != nil {
		fmt.Println(err)
	}

	return
}
