package repository

//go:generate mockgen -source=s3.go -destination=mock/s3.go
type S3 interface {
	SetBucketName(bucketName string)
	DownloadFile(key string, destinationPath string) error
	UploadFile(key string, filePath string) error
}
