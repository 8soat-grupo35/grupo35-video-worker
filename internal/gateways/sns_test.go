package gateways

import (
	"context"
	mock_wrappers "grupo35-video-worker/internal/adapters/wrappers/mock"
	"grupo35-video-worker/internal/entities"
	"grupo35-video-worker/internal/presenter"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/stretchr/testify/assert"

	"go.uber.org/mock/gomock"
)

func TestSendMessage_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	snsClient := mock_wrappers.NewMockISNSClient(ctrl)
	snsClient.EXPECT().Publish(context.TODO(), gomock.Any()).Return(&sns.PublishOutput{}, nil).AnyTimes()

	snsManager := NewSNS(snsClient, "topic-arn")
	err := snsManager.SendMessage(presenter.VideoStatus{
		User: entities.User{
			ID:    "id",
			Email: "email",
		},
		Status:  "status",
		ZipPath: "zip-path",
	})

	assert.NoError(t, err)
}

func TestSendMessage_Error_Marshal(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	snsClient := mock_wrappers.NewMockISNSClient(ctrl)

	snsManager := NewSNS(snsClient, "topic-arn")
	err := snsManager.SendMessage(make(chan int))

	assert.Error(t, err)
}

func TestSendMessage_Error_Publish(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	snsClient := mock_wrappers.NewMockISNSClient(ctrl)
	snsClient.EXPECT().Publish(context.TODO(), gomock.Any()).Return(nil, assert.AnError).AnyTimes()

	snsManager := NewSNS(snsClient, "topic-arn")
	err := snsManager.SendMessage(presenter.VideoStatus{
		User: entities.User{
			ID:    "id",
			Email: "email",
		},
		Status:  "status",
		ZipPath: "zip-path",
	})

	assert.Equal(t, err, assert.AnError)
}
