package gateways

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateZipWithScreenshots_Success(t *testing.T) {
	destinationPath := "test/data/test.zip"
	defer os.Remove(destinationPath)

	f, _ := os.Create("test/data/test.txt")
	f.Close()
	defer os.Remove("test/data/test.txt")

	zipGenerator := NewZipGenerator("")
	zipGenerator.CreateZipWithScreenshots(destinationPath, []string{"test/data/test.txt"})

	assert.FileExists(t, destinationPath)
}

func TestCreateZipWithScreenshots_Error_CreateZip(t *testing.T) {
	destinationPath := "test/data/"
	zipGenerator := NewZipGenerator("")

	err := zipGenerator.CreateZipWithScreenshots(destinationPath, []string{})

	assert.Error(t, err)
}
