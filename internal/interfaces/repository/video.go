package repository

//go:generate mockgen -source=video.go -destination=mock/video.go
type Video interface {
	SetVideoConfig(videoPath string, fileFormatPath string)
	GenerateVideoScreenshots(start float64, skipTime float64) (screenshotsPaths []string, err error)
}
