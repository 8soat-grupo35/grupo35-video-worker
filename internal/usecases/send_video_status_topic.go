package usecases

import (
	"fmt"
	"grupo35-video-worker/internal/gateways"
	"grupo35-video-worker/internal/presenter"
)

func SendVideoStatusTopic(sns gateways.SNS, videoStatus presenter.VideoStatus) error {
	fmt.Println("Sending status to SNS")
	return sns.SendMessage(videoStatus)
}
