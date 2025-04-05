package gateways

import (
	"context"
	"encoding/json"
	"grupo35-video-worker/internal/interfaces/repository"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type SNS struct {
	client   *sns.Client
	topicArn string
}

func NewSNS(cfg aws.Config, topicArn string) repository.SNS {
	return SNS{
		client:   sns.NewFromConfig(cfg),
		topicArn: topicArn,
	}
}

func (S SNS) SendMessage(message interface{}) error {
	convertedMessage, err := json.Marshal(message)

	if err != nil {
		return err
	}

	_, err = S.client.Publish(context.TODO(), &sns.PublishInput{
		TopicArn: aws.String(S.topicArn),
		Message:  aws.String(string(convertedMessage)),
	})

	return err
}
