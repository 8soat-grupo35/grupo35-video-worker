package main

import (
	"archive/zip"
	"io"
	"os"
)

type ZipGenerator struct {
	file      *os.File
	zipWriter *zip.Writer
}

func NewZipGenerator(destinationPath string) ZipGenerator {
	file, err := os.Create(destinationPath)

	if err != nil {
		panic(err)
	}

	w := zip.NewWriter(file)

	return ZipGenerator{
		file:      file,
		zipWriter: w,
	}
}

func (Z ZipGenerator) AddFiles(files []string) {
	for _, filePath := range files {

		fileToBeCreated, err := Z.zipWriter.Create(filePath)

		if err != nil {
			continue
		}

		file, err := os.Open(filePath)

		if err != nil {
			continue
		}

		io.Copy(fileToBeCreated, file)

		defer file.Close()
	}

	defer Z.file.Close()
}
