package usecases

import (
	"fmt"
	"grupo35-video-worker/internal/gateways"
)

//go:generate mockgen -source=create_zip_from_screenshots.go -destination=mock/create_zip_from_screenshots.go
func CreateZipFromScreenshots(screenshotsFiles []string) (zipPath string) {
	fmt.Println("Creating zip from screenshots")

	zipPath = "screenshots.zip"
	zip := gateways.NewZipGenerator("screenshots.zip")
	zip.AddFiles(screenshotsFiles)

	return zipPath
}
