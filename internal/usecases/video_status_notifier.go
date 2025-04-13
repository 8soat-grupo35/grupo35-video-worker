package usecases

import (
	"fmt"
	"grupo35-video-worker/internal/interfaces/repository"
	"grupo35-video-worker/internal/presenter"
)

//go:generate mockgen -source=video_status_notifier.go -destination=mock/video_status_notifier.go
type IVideoStatusNotifier interface {
	Notify(videoStatus presenter.VideoStatus) error
}

type VideoStatusNotifier struct {
	SNS repository.SNS
}

func NewVideoStatusNotifier(sns repository.SNS) IVideoStatusNotifier {
	return VideoStatusNotifier{
		SNS: sns,
	}
}

func (v VideoStatusNotifier) Notify(videoStatus presenter.VideoStatus) error {
	fmt.Println("Sending status to SNS")
	return v.SNS.SendMessage(videoStatus)
}
