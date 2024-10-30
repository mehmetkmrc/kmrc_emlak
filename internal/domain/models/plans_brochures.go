package models

import "github.com/Microsoft/go-winio/pkg/guid"

type (
	Plans_brochures struct {
		Property_id guid.GUID `json:"property_id"`
		File_type *string `json:"file_type"`
		File_path *string `json:"file_path"`
	}
)