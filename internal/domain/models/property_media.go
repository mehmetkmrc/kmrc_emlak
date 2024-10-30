package models

import "github.com/Microsoft/go-winio/pkg/guid"

type (
	Property_media struct {
		Property_id guid.GUID `json:"property_id" binding:"required"`
		Image []struct{
			ImageBase64 string `json:"image_base64"`
		} `json:"image"`
		Background_image []struct{
			ImageBase64 string `json:"image_base64"`
		} `json:"background_image"`
		Type string `json:"type" binding:"required"`
	}
)