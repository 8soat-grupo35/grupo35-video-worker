package controllers

import (
	"grupo35-video-worker/internal/adapters"
	"grupo35-video-worker/internal/interfaces/repository"
	"grupo35-video-worker/internal/usecases"
	"os"
)

type ProcessVideo struct {
	videoPathStructure adapters.VideoProcessPathStructure
	fileTransfer       usecases.ITransferFile
	videoProcessor     usecases.IProcessVideo
}

func NewProcessVideo(
	videoPathStructure adapters.VideoProcessPathStructure,
	s3manager repository.S3,
	videoProcessor repository.Video,
	zipProcessor repository.Zip,
) ProcessVideo {
	return ProcessVideo{
		videoPathStructure: videoPathStructure,
		fileTransfer:       usecases.NewTransferFile(s3manager),
		videoProcessor:     usecases.NewProcessVideo(videoProcessor, zipProcessor),
	}
}

func (p ProcessVideo) ProcessVideo() error {
	os.Mkdir(p.videoPathStructure.BasePath, 0777)
	os.Mkdir(p.videoPathStructure.BasePath+"/screenshots", 0777)
	defer os.RemoveAll(p.videoPathStructure.BasePath)

	_, err := p.fileTransfer.GetVideo(p.videoPathStructure.BucketKey, p.videoPathStructure.VideoOutputPath)

	if err != nil {
		return err
	}

	screenshotsFiles, err := p.videoProcessor.GenerateVideoScreenshots(p.videoPathStructure.VideoOutputPath, p.videoPathStructure.ScreenshotsOutputPath)

	if err != nil {
		return err
	}

	_, err = p.videoProcessor.CreateZipFromScreenshots(screenshotsFiles, p.videoPathStructure.ZipOutputPath)

	if err != nil {
		return err
	}

	err = p.fileTransfer.UploadZip(p.videoPathStructure.ZipOutputPath, p.videoPathStructure.ZipOutputPath)

	if err != nil {
		return err
	}

	return nil
}
