package controllers

import (
	"grupo35-video-worker/internal/adapters"
	"grupo35-video-worker/internal/entities"
	mock_repository "grupo35-video-worker/internal/interfaces/repository/mock"
	"grupo35-video-worker/internal/presenter"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNotify_Success(t *testing.T) {
	ctrl := gomock.NewController(t)

	videoToProcessMessage := adapters.VideoToProcess{
		User: entities.User{
			ID:    "1",
			Email: "teste@teste.com",
		},
		VideoPath: "video.mp4",
	}

	snsNotifier := mock_repository.NewMockSNS(ctrl)
	snsNotifier.EXPECT().SendMessage(presenter.VideoStatus{
		User:    videoToProcessMessage.User,
		Status:  "processed",
		ZipPath: "screenshots.zip",
	}).Return(nil).AnyTimes()

	notify := NewNotifyVideoStatus(videoToProcessMessage, snsNotifier)
	err := notify.Notify(true)

	assert.NoError(t, err)
}

func TestNotify_Success_Error_Message(t *testing.T) {
	ctrl := gomock.NewController(t)

	videoToProcessMessage := adapters.VideoToProcess{
		User: entities.User{
			ID:    "1",
			Email: "teste@teste.com",
		},
		VideoPath: "video.mp4",
	}

	snsNotifier := mock_repository.NewMockSNS(ctrl)
	snsNotifier.EXPECT().SendMessage(presenter.VideoStatus{
		User:    videoToProcessMessage.User,
		Status:  "error",
		ZipPath: "",
	}).Return(nil).AnyTimes()

	notify := NewNotifyVideoStatus(videoToProcessMessage, snsNotifier)
	err := notify.Notify(false)

	assert.NoError(t, err)
}

func TestNotify_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	videoToProcessMessage := adapters.VideoToProcess{
		User: entities.User{
			ID:    "1",
			Email: "teste@teste.com",
		},
		VideoPath: "video.mp4",
	}

	snsNotifier := mock_repository.NewMockSNS(ctrl)
	snsNotifier.EXPECT().SendMessage(presenter.VideoStatus{
		User:    videoToProcessMessage.User,
		Status:  "processed",
		ZipPath: "screenshots.zip",
	}).Return(assert.AnError).AnyTimes()

	notify := NewNotifyVideoStatus(videoToProcessMessage, snsNotifier)
	err := notify.Notify(true)

	assert.EqualError(t, err, assert.AnError.Error())
}
