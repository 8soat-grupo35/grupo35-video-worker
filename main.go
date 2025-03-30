package main

import (
	"context"
	"grupo35-video-worker/internal/handlers"

	"github.com/aws/aws-sdk-go-v2/config"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		panic(err)
	}

	handlers.ProcessVideos(cfg)
}
