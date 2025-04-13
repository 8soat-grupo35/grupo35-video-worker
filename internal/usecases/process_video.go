package usecases

import (
	"fmt"
	"grupo35-video-worker/internal/interfaces/repository"
)

//go:generate mockgen -source=process_video.go -destination=mock/process_video.go
type IProcessVideo interface {
	GenerateVideoScreenshots(videoPath string, screenshotsOutputPath string) (screenshotsFiles []string, err error)
	CreateZipFromScreenshots(screenshotsFiles []string, zipOutputPath string) (zipPath string, err error)
}

type ProcessVideo struct {
	VideoProcessor repository.Video
	ZipProcessor   repository.Zip
}

func NewProcessVideo(videoProcessor repository.Video, zipProcessor repository.Zip) IProcessVideo {
	return ProcessVideo{
		VideoProcessor: videoProcessor,
		ZipProcessor:   zipProcessor,
	}
}

func (p ProcessVideo) GenerateVideoScreenshots(videoPath string, screenshotsOutputPath string) (screenshotsFiles []string, err error) {
	fmt.Println("Processing video downloaded")
	p.VideoProcessor.SetVideoConfig(videoPath, screenshotsOutputPath)
	screenshotsFiles, err = p.VideoProcessor.GenerateVideoScreenshots(0, 1)

	return
}

func (p ProcessVideo) CreateZipFromScreenshots(screenshotsFiles []string, zipOutputPath string) (zipPath string, err error) {
	fmt.Println("Creating zip from screenshots")
	err = p.ZipProcessor.CreateZipWithScreenshots(zipOutputPath, screenshotsFiles)

	return zipOutputPath, err
}
