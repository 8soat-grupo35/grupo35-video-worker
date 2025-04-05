package usecases

import (
	"fmt"
	"grupo35-video-worker/internal/interfaces/repository"
)

func GetVideo(s3 repository.S3, videoPath string) (outputPath string, err error) {
	fmt.Println("Getting file from S3: ", videoPath)

	s3.SetBucketName("grupo35-video-uploaded")

	outputPath = "video.mp4"
	err = s3.DownloadFile(videoPath, outputPath)

	if err != nil {
		fmt.Println(err)
	}

	return
}
