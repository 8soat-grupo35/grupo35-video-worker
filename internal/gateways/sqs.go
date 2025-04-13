package gateways

import (
	"context"
	"fmt"
	"grupo35-video-worker/internal/adapters/wrappers"
	"grupo35-video-worker/internal/interfaces/repository"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type SQSHelper struct {
	client              wrappers.ISQSClient
	queueName           string
	maxNumberOfMessages int32
}

func NewSQSConsumer(client wrappers.ISQSClient, queueName string, maxNumberOfMessages int32) repository.SQS {
	return SQSHelper{
		client:              client,
		queueName:           queueName,
		maxNumberOfMessages: maxNumberOfMessages,
	}
}

func (S SQSHelper) ConsumeMessages(consumeFn func(message types.Message)) error {
	output, err := S.client.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(S.queueName),
		MaxNumberOfMessages: S.maxNumberOfMessages,
	})

	if err != nil {
		return err
	}

	fmt.Println("Getting messages from SQS length: ", len(output.Messages))

	for _, message := range output.Messages {
		consumeFn(message)

		_, err = S.client.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
			QueueUrl:      aws.String(S.queueName),
			ReceiptHandle: message.ReceiptHandle,
		})

		if err != nil {
			return err
		}
	}

	return nil
}
