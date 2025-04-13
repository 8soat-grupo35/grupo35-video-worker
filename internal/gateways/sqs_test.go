package gateways

import (
	"context"
	mock_wrappers "grupo35-video-worker/internal/adapters/wrappers/mock"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestConsumeMessage_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sqsClient := mock_wrappers.NewMockISQSClient(ctrl)
	sqsClient.EXPECT().ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String("queue-name"),
		MaxNumberOfMessages: 1,
	}).Return(&sqs.ReceiveMessageOutput{
		Messages: []types.Message{
			{
				Body:          aws.String("message"),
				ReceiptHandle: aws.String("receipt-handle"),
			},
		},
	}, nil).AnyTimes()
	sqsClient.EXPECT().DeleteMessage(gomock.Any(), gomock.Any()).Return(&sqs.DeleteMessageOutput{}, nil).AnyTimes()

	sqsHelper := NewSQSConsumer(sqsClient, "queue-name", 1)
	err := sqsHelper.ConsumeMessages(func(message types.Message) {
		// do nothing
	})

	assert.NoError(t, err)
}

func TestConsumeMessage_Error_ReceiveMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sqsClient := mock_wrappers.NewMockISQSClient(ctrl)
	sqsClient.EXPECT().ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String("queue-name"),
		MaxNumberOfMessages: 1,
	}).Return(nil, assert.AnError).AnyTimes()

	sqsHelper := NewSQSConsumer(sqsClient, "queue-name", 1)
	err := sqsHelper.ConsumeMessages(func(message types.Message) {
		// do nothing
	})
	assert.EqualError(t, err, assert.AnError.Error())
}

func TestConsumeMessage_Error_DeleteMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sqsClient := mock_wrappers.NewMockISQSClient(ctrl)
	sqsClient.EXPECT().ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String("queue-name"),
		MaxNumberOfMessages: 1,
	}).Return(&sqs.ReceiveMessageOutput{
		Messages: []types.Message{
			{
				Body:          aws.String("message"),
				ReceiptHandle: aws.String("receipt-handle"),
			},
		},
	}, nil).AnyTimes()

	sqsClient.EXPECT().DeleteMessage(gomock.Any(), gomock.Any()).Return(nil, assert.AnError).AnyTimes()

	sqsHelper := NewSQSConsumer(sqsClient, "queue-name", 1)
	err := sqsHelper.ConsumeMessages(func(message types.Message) {
		// do nothing
	})
	assert.EqualError(t, err, assert.AnError.Error())
}
