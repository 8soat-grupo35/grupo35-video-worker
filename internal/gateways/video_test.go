package gateways

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateVideoScreenshots_Success(t *testing.T) {
	os.Mkdir("test/data/screenshots", 0777)
	defer os.RemoveAll("test/data/screenshots")

	video := NewVideo("")
	video.SetVideoConfig("test/data/teste.mp4", "test/data/screenshots/test-%f.png")
	screenshots, err := video.GenerateVideoScreenshots(0, 1)

	assert.NoError(t, err)
	assert.Len(t, screenshots, 11)
}

func TestGenerateVideoScreenshots_Error_SetVideoConfig(t *testing.T) {
	video := NewVideo("")
	_, err := video.GenerateVideoScreenshots(0, 1)

	assert.EqualError(t, err, "set Video Config was not called")
}
