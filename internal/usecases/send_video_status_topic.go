package usecases

import (
	"fmt"
	"grupo35-video-worker/internal/interfaces/repository"
	"grupo35-video-worker/internal/presenter"
)

//go:generate mockgen -source=send_video_status_topic.go -destination=mock/send_video_status_topic.go
func SendVideoStatusTopic(sns repository.SNS, videoStatus presenter.VideoStatus) error {
	fmt.Println("Sending status to SNS")
	return sns.SendMessage(videoStatus)
}
