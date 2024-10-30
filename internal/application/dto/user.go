package dto

import "github.com/Microsoft/go-winio/pkg/guid"

type (
	UserRegisterRequest struct {
		First_name      string `json:"first_name" binding:"required"`
		Last_name       string `json:"last_name" binding:"required"`
		Email           string `json:"email" binding:"required"`
		Password        string `json:"password" binding:"required", min=8`
		ConfirmPassword string `binding:"required", min=8`
	}

	UserLoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,password"`
	}

	UserLoginResponse struct {
		User_id guid.GUID `json:"user_id"`
		First_name string `json:"first_name"`
		Last_name string `json:"last_name"`
		Role string `json:"role"`
		AccessToken string `json:"access_token"`
	}
)