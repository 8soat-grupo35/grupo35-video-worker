package adapters

import (
	"encoding/json"
	"grupo35-video-worker/internal/entities"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type VideoToProcess struct {
	User      entities.User `json:"user"`
	VideoPath string        `json:"video_path"`
}

func NewVideoToProcessFromSQSMessage(message types.Message) (*VideoToProcess, error) {
	var convertedMessage VideoToProcess

	err := json.Unmarshal([]byte(*message.Body), &convertedMessage)

	if err != nil {
		return nil, err
	}

	return &convertedMessage, nil
}
