package models

import (
	"time"

	"github.com/Microsoft/go-winio/pkg/guid"
)

type (
	User struct {
		User_id guid.GUID `json:"user_id" binding:"required"`
		First_name string `json:"first_name" binding:"required"`
		Last_name  string `json:"last_name" binding:"required"`
		Email      string `json:"email" binding:"required"`
		Phone string `json:"phone" binding:"required"`
		Password   string `json:"password" binding:"required"`
		Role string `json:"role" binding:"required"`
		Created_at time.Time `json:"created_at"`
		Updated_at time.Time `json:"updated_at"`
		Last_login time.Time `json:"last_login" binding:"required"`
	}

	
)