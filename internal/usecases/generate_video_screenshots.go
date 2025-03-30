package usecases

import (
	"fmt"
	"grupo35-video-worker/internal/gateways"
)

func GenerateVideoScreenshots(videoPath string) (screenshotsFiles []string) {
	fmt.Println("Processing video downloaded")
	video := gateways.NewVideo(videoPath, "screenshots/video_teste_output_%f.png")
	screenshotsFiles = video.GenerateVideoScreenshots(0, 1)

	return
}
