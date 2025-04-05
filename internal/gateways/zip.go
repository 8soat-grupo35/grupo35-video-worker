package gateways

import (
	"archive/zip"
	"fmt"
	"grupo35-video-worker/internal/interfaces/repository"
	"io"
	"os"
)

type ZipGenerator struct {
	destinationPath string
	file            *os.File
}

func NewZipGenerator(destinationPath string) repository.Zip {
	file, err := os.Create(destinationPath)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer file.Close()

	return ZipGenerator{
		destinationPath: destinationPath,
		file:            file,
	}
}

func (Z ZipGenerator) AddFiles(files []string) {
	zipFile, _ := os.OpenFile(Z.destinationPath, os.O_CREATE|os.O_WRONLY, 0644)
	defer zipFile.Close()

	w := zip.NewWriter(zipFile)
	defer w.Close()

	for _, filePath := range files {
		file, err := os.Open(filePath)

		if err != nil {
			fmt.Println("error open file", err)
			continue
		}
		defer file.Close()

		read, _ := io.ReadAll(file)
		fileToBeCreated, err := w.Create(file.Name())
		if err != nil {
			fmt.Println("error w.create", err)
			continue
		}

		fileToBeCreated.Write(read)
	}
}
