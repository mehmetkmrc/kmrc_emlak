package models

import "github.com/Microsoft/go-winio/pkg/guid"

type (
	Nearby struct {
		Property_id guid.GUID `json:"property_id" binding:"required"`
		Places *string `json:"places"`
		Distance *int `json:"distance"`
	}
)