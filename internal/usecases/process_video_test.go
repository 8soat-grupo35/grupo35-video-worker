package usecases

import (
	mock_repository "grupo35-video-worker/internal/interfaces/repository/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGenerateVideoScreenshots(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	screenshotsFiles := []string{"screenshots/video_teste_output_0.png"}

	videoProcessor := mock_repository.NewMockVideo(ctrl)
	videoProcessor.EXPECT().SetVideoConfig(gomock.Any(), gomock.Any()).Return().AnyTimes()
	videoProcessor.EXPECT().GenerateVideoScreenshots(gomock.Any(), gomock.Any()).Return(screenshotsFiles, nil).AnyTimes()

	zipProcessor := mock_repository.NewMockZip(ctrl)

	processVideo := ProcessVideo{
		VideoProcessor: videoProcessor,
		ZipProcessor:   zipProcessor,
	}

	response, err := processVideo.GenerateVideoScreenshots("video.mp4", "screenshots/video_teste_output_%f.png")

	assert.NoError(t, err)
	assert.Equal(t, screenshotsFiles, response)
}

func TestCreateZipFromScreenshots(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	screenshotsFiles := []string{"screenshots/video_teste_output_0.png"}
	screenshotsOutputPath := "screenshots.zip"

	videoProcessor := mock_repository.NewMockVideo(ctrl)
	zipProcessor := mock_repository.NewMockZip(ctrl)
	zipProcessor.EXPECT().CreateZipWithScreenshots("screenshots.zip", screenshotsFiles).Return(nil).AnyTimes()

	processVideo := ProcessVideo{
		VideoProcessor: videoProcessor,
		ZipProcessor:   zipProcessor,
	}

	response, err := processVideo.CreateZipFromScreenshots(screenshotsFiles, screenshotsOutputPath)

	assert.Equal(t, screenshotsOutputPath, response)
	assert.NoError(t, err)
}
