package gateways

import (
	"context"
	"encoding/json"
	"grupo35-video-worker/internal/adapters/wrappers"
	"grupo35-video-worker/internal/interfaces/repository"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type SNS struct {
	Client   wrappers.ISNSClient
	TopicArn string
}

func NewSNS(client wrappers.ISNSClient, topicArn string) repository.SNS {
	return SNS{
		Client:   client,
		TopicArn: topicArn,
	}
}

func (S SNS) SendMessage(message interface{}) error {
	convertedMessage, err := json.Marshal(message)

	if err != nil {
		return err
	}

	_, err = S.Client.Publish(context.TODO(), &sns.PublishInput{
		TopicArn: aws.String(S.TopicArn),
		Message:  aws.String(string(convertedMessage)),
	})

	return err
}
