package handlers

import (
	"fmt"
	"grupo35-video-worker/internal/adapters"
	"grupo35-video-worker/internal/controllers"
	"grupo35-video-worker/internal/interfaces/repository"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

//go:generate mockgen -source=process_videos.go -destination=mock/process_videos.go
type IProcessVideosHandler interface {
	ProcessVideos()
}
type ProcessVideosConfig struct {
	SQS   repository.SQS
	SNS   repository.SNS
	S3    repository.S3
	Video repository.Video
	Zip   repository.Zip
}

func NewProcessVideosHandler(config ProcessVideosConfig) IProcessVideosHandler {
	return config
}

func (p ProcessVideosConfig) ProcessVideos() {
	for {
		p.SQS.ConsumeMessages(func(message types.Message) {
			fmt.Println(message)
			videoToProcess, err := adapters.NewVideoToProcessFromSQSMessage(message)

			if err != nil {
				fmt.Println(err)
				return
			}

			videoPathStructure := adapters.GetVideoProcessPathStructure(*videoToProcess)

			controller := controllers.NewProcessVideo(videoPathStructure, p.S3, p.Video, p.Zip)
			err = controller.ProcessVideo()

			notifier := controllers.NewNotifyVideoStatus(
				*videoToProcess,
				p.SNS,
			)

			err = notifier.Notify(err == nil, videoPathStructure.ZipOutputPath)

			if err != nil {
				fmt.Println(err)
			}
		})

	}
}
