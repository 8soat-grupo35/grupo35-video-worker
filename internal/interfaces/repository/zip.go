package repository

type Zip interface {
	AddFiles(files []string)
}
