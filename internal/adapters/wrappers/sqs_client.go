package wrappers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

//go:generate mockgen -source=sqs_client.go -destination=mock/sqs_client.go
type ISQSClient interface {
	ReceiveMessage(ctx context.Context, params *sqs.ReceiveMessageInput, optFns ...func(*sqs.Options)) (*sqs.ReceiveMessageOutput, error)
	DeleteMessage(ctx context.Context, params *sqs.DeleteMessageInput, optFns ...func(*sqs.Options)) (*sqs.DeleteMessageOutput, error)
}

type SQSClient struct {
	client *sqs.Client
}

// ReceiveMessage implements ISQSClient.
func (s *SQSClient) ReceiveMessage(ctx context.Context, params *sqs.ReceiveMessageInput, optFns ...func(*sqs.Options)) (*sqs.ReceiveMessageOutput, error) {
	return s.client.ReceiveMessage(ctx, params, optFns...)
}

// DeleteMessage implements ISQSClient.
func (s *SQSClient) DeleteMessage(ctx context.Context, params *sqs.DeleteMessageInput, optFns ...func(*sqs.Options)) (*sqs.DeleteMessageOutput, error) {
	return s.client.DeleteMessage(ctx, params, optFns...)
}

func NewSQSClient(cfg aws.Config) ISQSClient {
	return &SQSClient{
		client: sqs.NewFromConfig(cfg),
	}
}
