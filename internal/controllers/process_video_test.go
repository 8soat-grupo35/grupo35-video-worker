package controllers

import (
	"grupo35-video-worker/internal/adapters"
	"grupo35-video-worker/internal/entities"
	mock_usecases "grupo35-video-worker/internal/usecases/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestProcessVideo_Sucess(t *testing.T) {
	ctrl := gomock.NewController(t)
	videoPath := "video.mp4"

	videoPathStructure := adapters.GetVideoProcessPathStructure(adapters.VideoToProcess{
		User: entities.User{
			ID:    "1",
			Email: "teste@teste.com",
		},
		VideoPath: videoPath,
	})

	transferFiles := mock_usecases.NewMockITransferFile(ctrl)
	transferFiles.EXPECT().GetVideo(videoPath, gomock.Any()).Return(videoPath, nil).AnyTimes()
	transferFiles.EXPECT().UploadZip(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	videoProcessor := mock_usecases.NewMockIProcessVideo(ctrl)
	videoProcessor.EXPECT().GenerateVideoScreenshots(gomock.Any(), gomock.Any()).Return([]string{"screenshots/video_teste_output_0.png"}, nil).AnyTimes()
	videoProcessor.EXPECT().CreateZipFromScreenshots(gomock.Any(), gomock.Any()).Return("screenshots.zip", nil).AnyTimes()

	processVideo := ProcessVideo{
		videoPathStructure: videoPathStructure,
		fileTransfer:       transferFiles,
		videoProcessor:     videoProcessor,
	}

	err := processVideo.ProcessVideo()

	assert.NoError(t, err)
}

func TestProcessVideo_Error_GetVideo(t *testing.T) {
	ctrl := gomock.NewController(t)
	videoPath := "video.mp4"

	videoPathStructure := adapters.GetVideoProcessPathStructure(adapters.VideoToProcess{
		User: entities.User{
			ID:    "1",
			Email: "teste@teste.com",
		},
		VideoPath: videoPath,
	})

	transferFiles := mock_usecases.NewMockITransferFile(ctrl)
	transferFiles.EXPECT().GetVideo(videoPath, gomock.Any()).Return("", assert.AnError).AnyTimes()

	videoProcessor := mock_usecases.NewMockIProcessVideo(ctrl)

	processVideo := ProcessVideo{
		videoPathStructure: videoPathStructure,
		fileTransfer:       transferFiles,
		videoProcessor:     videoProcessor,
	}

	err := processVideo.ProcessVideo()

	assert.EqualError(t, err, assert.AnError.Error())
}

func TestProcessVideo_Error_UploadZip(t *testing.T) {
	ctrl := gomock.NewController(t)
	videoPath := "video.mp4"
	videoPathStructure := adapters.GetVideoProcessPathStructure(adapters.VideoToProcess{
		User: entities.User{
			ID:    "1",
			Email: "teste@teste.com",
		},
		VideoPath: videoPath,
	})

	transferFiles := mock_usecases.NewMockITransferFile(ctrl)
	transferFiles.EXPECT().GetVideo(videoPath, gomock.Any()).Return(videoPath, nil).AnyTimes()
	transferFiles.EXPECT().UploadZip(gomock.Any(), gomock.Any()).Return(assert.AnError).AnyTimes()

	videoProcessor := mock_usecases.NewMockIProcessVideo(ctrl)
	videoProcessor.EXPECT().GenerateVideoScreenshots(gomock.Any(), gomock.Any()).Return([]string{"screenshots/video_teste_output_0.png"}, nil).AnyTimes()
	videoProcessor.EXPECT().CreateZipFromScreenshots(gomock.Any(), gomock.Any()).Return("screenshots.zip", nil).AnyTimes()

	processVideo := ProcessVideo{
		videoPathStructure: videoPathStructure,
		fileTransfer:       transferFiles,
		videoProcessor:     videoProcessor,
	}

	err := processVideo.ProcessVideo()
	assert.EqualError(t, err, assert.AnError.Error())
}
