package models

import "github.com/Microsoft/go-winio/pkg/guid"

type (
	Basic_info struct {
		Property_id guid.GUID `json:"property_id" binding:"required"`
		Basic_info_id guid.GUID `json:"basic_info_id" binding:"required"`
		Main_title string `json:"main_title" binding:"required"`
		Type string `json:"type" binding:"required"`
		Category string `json:"category" binding:"required"`
		Price int `json:"price" binding:"required"`
		Keywords string `json:"keywords" binding:"required"`
	}
)