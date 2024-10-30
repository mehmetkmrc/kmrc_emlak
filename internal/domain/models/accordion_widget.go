package models

import "github.com/Microsoft/go-winio/pkg/guid"

type (
	Accordion_widget struct {
		Property_id guid.GUID `json:"property_id" binding:"required"`
		Accordion_exist bool `json:"accordion_exist" binding:"required"`
		Accordion_title string `json:"accordion_title" binding:"required"`
		Accordion_details string `json:"accordion_details" binding:"required"`
	}
)
