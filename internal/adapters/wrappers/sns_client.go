package wrappers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

//go:generate mockgen -source=sns_client.go -destination=mock/sns_client.go
type ISNSClient interface {
	Publish(ctx context.Context, params *sns.PublishInput, optFns ...func(*sns.Options)) (*sns.PublishOutput, error)
}

type SNSClient struct {
	client *sns.Client
}

// Publish implements ISNSClient.
func (s *SNSClient) Publish(ctx context.Context, params *sns.PublishInput, optFns ...func(*sns.Options)) (*sns.PublishOutput, error) {
	return s.client.Publish(ctx, params, optFns...)
}

func NewSNSClient(cfg aws.Config) ISNSClient {
	return &SNSClient{
		client: sns.NewFromConfig(cfg),
	}
}
