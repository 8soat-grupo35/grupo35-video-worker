package adapters

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type VideoProcessPathStructure struct {
	BucketKey             string
	FileName              string
	FilePath              string
	FileExtension         string
	BasePath              string
	VideoOutputPath       string
	ScreenshotsOutputPath string
	ZipOutputPath         string
}

func GetVideoProcessPathStructure(videoToProcess VideoToProcess) VideoProcessPathStructure {
	videoPathStructure := VideoProcessPathStructure{}
	videoPathStructure.BucketKey = videoToProcess.VideoPath

	fmt.Println("video path", videoToProcess.VideoPath)
	folders := strings.Split(videoToProcess.VideoPath, "/")
	fmt.Println("folders", folders)
	fileName := folders[len(folders)-1]
	fmt.Println("fileName", fileName)

	videoPathStructure.FilePath = strings.Join(folders[:len(folders)-1], "/")
	videoPathStructure.FileName = strings.Split(fileName, ".")[0]
	videoPathStructure.FileExtension = strings.Split(fileName, ".")[1]

	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	videoPathStructure.BasePath = "video/" + timestamp + videoToProcess.User.ID

	videoPathStructure.VideoOutputPath = videoPathStructure.BasePath + "/" + fileName
	videoPathStructure.ScreenshotsOutputPath = videoPathStructure.BasePath + "/screenshots/screenshot_%f.png"
	videoPathStructure.ZipOutputPath = videoPathStructure.BasePath + "/" + videoPathStructure.FileName + ".zip"

	fmt.Println("videoPathStructure", videoPathStructure)

	return videoPathStructure
}
