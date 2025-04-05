package repository

import "github.com/aws/aws-sdk-go-v2/service/sqs/types"

type SQS interface {
	ConsumeMessages(consumeFn func(message types.Message))
}
