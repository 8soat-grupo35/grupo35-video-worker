package handlers

import (
	"fmt"
	"grupo35-video-worker/internal/adapters"
	"grupo35-video-worker/internal/adapters/wrappers"
	"grupo35-video-worker/internal/controllers"
	"grupo35-video-worker/internal/gateways"
	"grupo35-video-worker/internal/presenter"
	"grupo35-video-worker/internal/usecases"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

func ProcessVideos(cfg aws.Config) {
	sqsClient := wrappers.NewSQSClient(cfg)
	snsClient := wrappers.NewSNSClient(cfg)
	s3Client := wrappers.NewS3Client(cfg)

	for {

		consumer := gateways.NewSQSConsumer(sqsClient, "video-process-queue", 10)

		consumer.ConsumeMessages(func(message types.Message) {
			videoToProcess, err := adapters.NewVideoToProcessFromSQSMessage(message)

			if err != nil {
				fmt.Println(err)
				return
			}

			err = controllers.ProcessVideo(s3Client, videoToProcess.VideoPath)

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
			snsClient := gateways.NewSNS(snsClient, "arn:aws:sns:us-east-1:633053670772:video-status-topic")
			err = usecases.SendVideoStatusTopic(snsClient, snsResponse)

			if err != nil {
				fmt.Println(err)
			}
		})

	}
}
