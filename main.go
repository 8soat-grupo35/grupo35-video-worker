package main

import (
	"os"
)

func main() {
	os.Mkdir("screenshots", 0777)
	video := NewVideo("video_teste.mp4", "screenshots/video_teste_output_%f.png")
	screenshotsFiles := video.GenerateVideoScreenshots(0, 1)

	zip := NewZipGenerator("screenshots.zip")
	zip.AddFiles(screenshotsFiles)
	os.RemoveAll("screenshots")
}
