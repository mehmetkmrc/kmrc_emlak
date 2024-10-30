package models

import "github.com/Microsoft/go-winio/pkg/guid"

type (
	Amenities struct {
		Property_id guid.GUID `json:"property_id" binding:"required"`
		Wifi *bool `json:"wifi" binding:"required"`
		Pool *bool `json:"pool" binding:"required"`
		Security *bool `json:"security" binding:"required"`
		Laundry_room *bool `json:"laundry_room" binding:"required"`
		Equipped_kitchen *bool `json:"equipped_kitchen" binding:"required"`
		Air_conditioning *bool `json:"air_conditioning" binding:"required"`
		Parking *bool `json:"parking" binding:"required"`
		Garage_atached *bool `json:"garage_atached" binding:"required"`
		Fireplace *bool `json:"fireplace" binding:"required"`
		Window_covering *bool `json:"window_covering" binding:"required"`
		Backyard *bool `json:"backyard" binding:"required"`
		Fitness_gym *bool `json:"fitness_gym" binding:"required"`
		Elevator *bool `json:"elevator" binding:"required"`
	}
)