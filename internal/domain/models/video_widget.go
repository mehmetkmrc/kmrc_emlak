package models

import "github.com/Microsoft/go-winio/pkg/guid"

type (
	Video_widget struct {
		Property_id guid.GUID `json:"property_id"`
		Video_exist *bool `json:"video_exist"`
		Video_title *string `json:"video_title"`
		Youtube_url *string `json:"youtube_url"`
		Vimeo_url *string `json:"vimeo_url"`
	}
)