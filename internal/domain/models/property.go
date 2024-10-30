package models

import "github.com/Microsoft/go-winio/pkg/guid"

type (
	Property struct {
		User_id guid.GUID `json:"user_id" binding:"required"`
		Property_id guid.GUID `json:"property_id" binding:"required"`
		Tariff_plan string `json:"tariff_plan" binding:"required"`
	}
)