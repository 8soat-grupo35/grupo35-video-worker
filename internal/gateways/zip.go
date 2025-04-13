package gateways

import (
	"archive/zip"
	"fmt"
	"grupo35-video-worker/internal/interfaces/repository"
	"io"
	"os"
)

type ZipGenerator struct {
	basePath string
}

func NewZipGenerator(basePath string) repository.Zip {
	return &ZipGenerator{
		basePath: basePath,
	}
}

func (Z *ZipGenerator) createZip(destinationPath string) error {
	file, err := os.Create(destinationPath)

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer file.Close()

	return nil
}

func (Z *ZipGenerator) CreateZipWithScreenshots(destinationPath string, files []string) error {
	err := Z.createZip(Z.basePath + destinationPath)

	if err != nil {
		return err
	}

	zipFile, _ := os.OpenFile(Z.basePath+destinationPath, os.O_CREATE|os.O_WRONLY, 0644)
	defer zipFile.Close()

	w := zip.NewWriter(zipFile)
	defer w.Close()

	for _, filePath := range files {
		file, err := os.Open(filePath)

		if err != nil {
			fmt.Println("error open file", err)
			return err
		}
		defer file.Close()

		read, _ := io.ReadAll(file)
		fileToBeCreated, err := w.Create(file.Name())
		if err != nil {
			fmt.Println("error w.create", err)
			return err
		}

		fileToBeCreated.Write(read)
	}

	return nil
}
