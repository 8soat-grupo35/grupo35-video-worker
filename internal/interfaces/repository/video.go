package repository

type Video interface {
	GenerateVideoScreenshots(start float64, skipTime float64) (screenshotsPaths []string)
}
