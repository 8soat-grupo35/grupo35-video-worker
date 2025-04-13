package usecases

import (
	"fmt"
	"grupo35-video-worker/internal/interfaces/repository"
)

//go:generate mockgen -source=process_video.go -destination=mock/process_video.go
type IProcessVideo interface {
	GenerateVideoScreenshots(videoPath string) (screenshotsFiles []string, err error)
	CreateZipFromScreenshots(screenshotsFiles []string) (zipPath string, err error)
}

type ProcessVideoConfig struct {
	VideoPath         string
	ScreenshotsOutput string
	ZipPath           string
}

type ProcessVideo struct {
	ProcessVideoConfig
	VideoProcessor repository.Video
	ZipProcessor   repository.Zip
}

func NewProcessVideo(processVideo ProcessVideo) IProcessVideo {
	return processVideo
}

func (p ProcessVideo) GenerateVideoScreenshots(videoPath string) (screenshotsFiles []string, err error) {
	fmt.Println("Processing video downloaded")
	p.VideoProcessor.SetVideoConfig(videoPath, p.ScreenshotsOutput)
	screenshotsFiles, err = p.VideoProcessor.GenerateVideoScreenshots(0, 1)

	return
}

func (p ProcessVideo) CreateZipFromScreenshots(screenshotsFiles []string) (zipPath string, err error) {
	fmt.Println("Creating zip from screenshots")
	err = p.ZipProcessor.CreateZipWithScreenshots(p.ZipPath, screenshotsFiles)

	return p.ZipPath, err
}
