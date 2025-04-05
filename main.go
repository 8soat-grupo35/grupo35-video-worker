package main

import (
	"context"
	"fmt"
	"grupo35-video-worker/internal/handlers"

	"github.com/aws/aws-sdk-go-v2/config"
)

func main() {
	fmt.Println("Starting video worker")
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	handlers.ProcessVideos(cfg)
}
