package presenter

import "grupo35-video-worker/internal/entities"

type VideoStatus struct {
	User    entities.User
	Status  string `json:"status"`
	ZipPath string `json:"zip_path"`
	VideoPath string `json:"video_path"`
}
