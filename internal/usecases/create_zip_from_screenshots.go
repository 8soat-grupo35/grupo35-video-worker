package usecases

import (
	"fmt"
	"grupo35-video-worker/internal/gateways"
)

func CreateZipFromScreenshots(screenshotsFiles []string) (zipPath string) {
	fmt.Println("Creating zip from screenshots")

	zipPath = "screenshots.zip"
	zip := gateways.NewZipGenerator("screenshots.zip")
	zip.AddFiles(screenshotsFiles)

	return zipPath
}
