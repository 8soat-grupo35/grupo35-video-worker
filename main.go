package main

import (
	"archive/zip"
	"fmt"
	"os"

	"github.com/mowshon/moviego"
)

func main() {

	video := NewVideo("video_teste.mp4", "video_teste_output_%f.png")
	screenshotsFiles := video.GenerateVideoScreenshots(0, 1)

	zip := NewZipGenerator("screenshots.zip")
	zip.AddFiles(screenshotsFiles)
}

