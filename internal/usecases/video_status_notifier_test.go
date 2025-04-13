package usecases

import (
	"grupo35-video-worker/internal/entities"
	mock_repository "grupo35-video-worker/internal/interfaces/repository/mock"
	"grupo35-video-worker/internal/presenter"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNotify_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	videoStatus := presenter.VideoStatus{
		User: entities.User{
			ID:    "1",
			Email: "email@email.com",
		},
		Status:  "success",
		ZipPath: "zip_path.zip",
	}

	snsMock := mock_repository.NewMockSNS(ctrl)

	snsMock.EXPECT().SendMessage(videoStatus).Return(nil).AnyTimes()

	videoStatusNotifier := NewVideoStatusNotifier(snsMock)

	err := videoStatusNotifier.Notify(videoStatus)

	assert.NoError(t, err)
}

func TestNotify_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	videoStatus := presenter.VideoStatus{
		User: entities.User{
			ID:    "1",
			Email: "email@email.com",
		},
		Status:  "success",
		ZipPath: "zip_path.zip",
	}
	snsMock := mock_repository.NewMockSNS(ctrl)
	snsMock.EXPECT().SendMessage(videoStatus).Return(assert.AnError).AnyTimes()

	videoStatusNotifier := NewVideoStatusNotifier(snsMock)

	err := videoStatusNotifier.Notify(videoStatus)
	assert.EqualError(t, err, assert.AnError.Error())
}
