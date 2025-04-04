package usecases

import (
	"fmt"
	"grupo35-video-worker/internal/interfaces/repository"
)

func UploadZip(s3 repository.S3, zipPath string) (err error) {
	fmt.Println("Uploading processed video to S3")
	s3.SetBucketName("grupo35-video-processed")
	err = s3.UploadFile("screenshots.zip", "screenshots.zip")

	if err != nil {
		fmt.Println(err)
	}

	return
}
