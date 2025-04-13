package repository

//go:generate mockgen -source=zip.go -destination=mock/zip.go
type Zip interface {
	CreateZipWithScreenshots(destinationFiles string, files []string) error
}
