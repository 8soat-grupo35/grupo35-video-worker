package main

import (
	"context"
	"fmt"
	"grupo35-video-worker/internal/adapters/wrappers"
	"grupo35-video-worker/internal/gateways"
	"grupo35-video-worker/internal/handlers"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
)

func main() {
	fmt.Println("Starting video worker")
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("AWS Config loaded")

	os.Mkdir("video/", 0777)

	sqsClient := wrappers.NewSQSClient(cfg)
	snsClient := wrappers.NewSNSClient(cfg)
	s3Client := wrappers.NewS3Client(cfg)

	videoProcessorHandler := handlers.NewProcessVideosHandler(handlers.ProcessVideosConfig{
		SQS:   gateways.NewSQSConsumer(sqsClient, "video-process-queue", 10),
		SNS:   gateways.NewSNS(snsClient, "arn:aws:sns:us-east-1:"+os.Getenv("AWS_ACCOUNT_ID")+":video-status-topic"),
		S3:    gateways.NewS3Manager(s3Client),
		Video: gateways.NewVideo(),
		Zip:   gateways.NewZipGenerator(),
	})

	videoProcessorHandler.ProcessVideos()
}
