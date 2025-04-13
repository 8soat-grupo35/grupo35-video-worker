package gateways

import (
	"errors"
	"fmt"
	"grupo35-video-worker/internal/interfaces/repository"

	"github.com/mowshon/moviego"
)

type Video struct {
	basePath       string
	videoPath      *string
	fileFormatPath *string
}

func NewVideo(basePath string) repository.Video {
	return &Video{
		basePath: basePath,
	}
}

func (V *Video) SetVideoConfig(videoPath string, fileFormatPath string) {
	V.videoPath = &videoPath
	V.fileFormatPath = &fileFormatPath
}

func (V *Video) GenerateVideoScreenshots(start float64, skipTime float64) ([]string, error) {
	if V.videoPath == nil || V.fileFormatPath == nil {
		return []string{}, errors.New("set Video Config was not called")
	}

	video, _ := moviego.Load(V.basePath + *V.videoPath)
	screenshots := []string{}

	for i := start; i < video.Duration(); i += skipTime {
		fmt.Println(fmt.Sprintf(*V.fileFormatPath, i))
		video.Screenshot(i, fmt.Sprintf(*V.fileFormatPath, i))
		screenshots = append(screenshots, fmt.Sprintf(*V.fileFormatPath, i))
	}

	return screenshots, nil
}
