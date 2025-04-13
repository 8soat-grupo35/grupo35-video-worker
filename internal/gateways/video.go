package gateways

import (
	"errors"
	"fmt"
	"grupo35-video-worker/internal/interfaces/repository"

	"github.com/mowshon/moviego"
)

type Video struct {
	videoPath      *string
	fileFormatPath *string
}

func NewVideo() repository.Video {
	return &Video{}
}

func (V *Video) SetVideoConfig(videoPath string, fileFormatPath string) {
	V.videoPath = &videoPath
	V.fileFormatPath = &fileFormatPath
}

func (V *Video) GenerateVideoScreenshots(start float64, skipTime float64) ([]string, error) {
	if V.videoPath == nil || V.fileFormatPath == nil {
		return []string{}, errors.New("set Video Config was not called")
	}

	fmt.Println("Generating screenshots from", *V.videoPath)

	video, _ := moviego.Load(*V.videoPath)
	screenshots := []string{}

	for i := start; i < video.Duration(); i += skipTime {
		fmt.Println("Generating screenshot at", fmt.Sprintf(*V.fileFormatPath, i))
		_, err := video.Screenshot(i, fmt.Sprintf(*V.fileFormatPath, i))

		if err == nil {
			screenshots = append(screenshots, fmt.Sprintf(*V.fileFormatPath, i))
		}
	}

	return screenshots, nil
}
