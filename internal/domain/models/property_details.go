package models

import "github.com/Microsoft/go-winio/pkg/guid"

type (
	Property_details struct {
		Property_id guid.GUID `json:"property_id" binding:"required"`
		Area float32 `json:"area" binding:"required"`
		Bedrooms *int `json:"bedrooms" binding:"required"`
		Bathrooms *int `json:"bathrooms" biding:"required"`
		Parking *int `json:"parking" binding:"required"`
		Accomodation *string `json:"accomodation" binding:"required"`
		Website *string `json:"website"`
		Property_details *string `json:"property_details"`
	}
)