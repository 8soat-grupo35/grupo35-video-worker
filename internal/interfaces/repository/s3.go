package repository

type S3 interface {
	SetBucketName(bucketName string)
	DownloadFile(key string, destinationPath string) error
	UploadFile(key string, filePath string) error
}
