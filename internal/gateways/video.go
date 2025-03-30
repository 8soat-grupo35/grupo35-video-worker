package gateways

import (
	"fmt"

	"github.com/mowshon/moviego"
)

type Video struct {
	video          moviego.Video
	fileFormatPath string
}

func NewVideo(videoPath string, fileFormatPath string) Video {
	video, _ := moviego.Load(videoPath)

	return Video{
		video:          video,
		fileFormatPath: fileFormatPath,
	}
}

func (V Video) GenerateVideoScreenshots(start float64, skipTime float64) []string {
	screenshots := []string{}

	for i := start; i < V.video.Duration(); i += skipTime {
		fmt.Println(fmt.Sprintf(V.fileFormatPath, i))
		V.video.Screenshot(i, fmt.Sprintf(V.fileFormatPath, i))
		screenshots = append(screenshots, fmt.Sprintf(V.fileFormatPath, i))
	}

	return screenshots
}
