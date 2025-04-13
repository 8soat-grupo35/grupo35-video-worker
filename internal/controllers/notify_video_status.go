package controllers

import (
	"fmt"
	"grupo35-video-worker/internal/adapters"
	"grupo35-video-worker/internal/interfaces/repository"
	"grupo35-video-worker/internal/presenter"
	"grupo35-video-worker/internal/usecases"
)

type NotifyVideoStatus struct {
	videoToProcessMessage adapters.VideoToProcess
	SNSNotifier           repository.SNS
}

func NewNotifyVideoStatus(videoToProcessMessage adapters.VideoToProcess, snsNotifier repository.SNS) NotifyVideoStatus {
	return NotifyVideoStatus{
		videoToProcessMessage: videoToProcessMessage,
		SNSNotifier:           snsNotifier,
	}
}

func (n NotifyVideoStatus) Notify(success bool) error {
	snsResponse := presenter.VideoStatus{
		User:    n.videoToProcessMessage.User,
		Status:  "processed",
		ZipPath: "screenshots.zip",
	}

	if !success {
		snsResponse.Status = "error"
		snsResponse.ZipPath = ""
	}

	fmt.Println("Sending status to SNS")
	videoStatusNotifier := usecases.NewVideoStatusNotifier(n.SNSNotifier)

	return videoStatusNotifier.Notify(snsResponse)
}
