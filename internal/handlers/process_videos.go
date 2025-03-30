package handlers

import (
	"fmt"
	"grupo35-video-worker/internal/adapters"
	"grupo35-video-worker/internal/controllers"
	"grupo35-video-worker/internal/gateways"
	"grupo35-video-worker/internal/presenter"
	"grupo35-video-worker/internal/usecases"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

func ProcessVideos(cfg aws.Config) {
	for {
		consumer := gateways.NewSQSConsumer(cfg, "video-process-queue", 10)

		consumer.ConsumeMessages(func(message types.Message) {
			videoToProcess, err := adapters.NewVideoToProcessFromSQSMessage(message)

			if err != nil {
				fmt.Println(err)
				return
			}

			err = controllers.ProcessVideo(cfg, videoToProcess.VideoPath)

			snsResponse := presenter.VideoStatus{
				User:    videoToProcess.User,
				Status:  "processed",
				ZipPath: "screenshots.zip",
			}

			if err != nil {
				snsResponse.Status = "error"
				snsResponse.ZipPath = ""
			}

			fmt.Println("Sending status to SNS")
			snsClient := gateways.NewSNS(cfg, "arn:aws:sns:us-east-1:633053670772:video-status-topic")
			err = usecases.SendVideoStatusTopic(snsClient, snsResponse)

			if err != nil {
				fmt.Println(err)
			}
		})

	}
}
