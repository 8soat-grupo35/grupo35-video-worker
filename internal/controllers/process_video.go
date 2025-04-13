package controllers

import (
	"grupo35-video-worker/internal/interfaces/repository"
	"grupo35-video-worker/internal/usecases"
	"os"
)

type ProcessVideo struct {
	videoPath      string
	fileTransfer   usecases.ITransferFile
	videoProcessor usecases.IProcessVideo
}

func NewProcessVideo(videoPath string, s3manager repository.S3, videoProcessor repository.Video, zipProcessor repository.Zip) ProcessVideo {

	return ProcessVideo{
		videoPath:    videoPath,
		fileTransfer: usecases.NewTransferFile(s3manager),
		videoProcessor: usecases.NewProcessVideo(usecases.ProcessVideo{
			ProcessVideoConfig: usecases.ProcessVideoConfig{
				VideoPath:         videoPath,
				ScreenshotsOutput: "screenshots/video_teste_output_%f.png",
				ZipPath:           "screenshots.zip",
			},
			VideoProcessor: videoProcessor,
			ZipProcessor:   zipProcessor,
		}),
	}
}

func (p ProcessVideo) ProcessVideo() error {
	os.Mkdir("screenshots", 0777)
	defer os.RemoveAll("screenshots")

	output := "video.mp4"

	outputVideo, err := p.fileTransfer.GetVideo(p.videoPath, output)

	if err != nil {
		return err
	}

	screenshotsFiles, err := p.videoProcessor.GenerateVideoScreenshots(outputVideo)

	if err != nil {
		return err
	}

	zipPath, err := p.videoProcessor.CreateZipFromScreenshots(screenshotsFiles)

	if err != nil {
		return err
	}

	err = p.fileTransfer.UploadZip(zipPath)

	if err != nil {
		return err
	}

	return nil
}
