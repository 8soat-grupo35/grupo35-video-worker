package repository

//go:generate mockgen -source=zip.go -destination=mock/zip.go
type Zip interface {
	AddFiles(files []string)
}
