package models

import "github.com/Microsoft/go-winio/pkg/guid"

type (
	Location struct {
		Property_id guid.GUID `json:"property_id" binding:"required"`
		Location_id guid.GUID `json:"location_id" binding:"required"`
		Phone string `json:"phone" binding:"required"`
		Mail string `json:"mail" binding:"required"`
		City string `json:"city" binding:"required"`
		Address string `json:"address" binding:"required"`
		Longitude float32 `json:"longitude" binding:"required"`
		Latitude float32 `json:"latitude" binding:"required"`
	}
)