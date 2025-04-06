package controllers

import (
	"grupo35-video-worker/internal/adapters/wrappers"
	"grupo35-video-worker/internal/gateways"
	"grupo35-video-worker/internal/usecases"
	"os"
)

func ProcessVideo(s3Client wrappers.IS3Client, videoPath string) error {
	os.Mkdir("screenshots", 0777)
	defer os.RemoveAll("screenshots")

	s3 := gateways.NewS3Manager(s3Client)
	outputVideo, err := usecases.GetVideo(s3, videoPath)

	if err != nil {
		return err
	}

	screenshotsFiles := usecases.GenerateVideoScreenshots(outputVideo)

	zipPath := usecases.CreateZipFromScreenshots(screenshotsFiles)

	err = usecases.UploadZip(s3, zipPath)

	if err != nil {
		return err
	}

	return nil
}
