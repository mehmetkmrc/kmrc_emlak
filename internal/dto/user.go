package dto

import "time"

type (
	// Requests
	UserLoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,password"`
	}

	// Responses
	UserDetail struct {
		ID      string `json:"id"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Surname string `json:"surname"`
	}

	GetUserResponse struct {
		ID          string    `json:"id"`
		Name        string    `json:"name"`
		Surname     string    `json:"surname"`
		Email       string    `json:"email"`
		CreatedAt   time.Time `json:"created_at"`
	}

	UserRegisterRequest struct{
		Name        string    `json:"name" validate:"required"`
		Surname     string    `json:"surname" validate:"required"`
		Email       string    `json:"email" validate:"required"`
		Phone		string	  `json:"phone" validate:"required"`
		Password 	string	  `json:"password" validate:"required"`
		ConfirmPassword string `json:"confirm_password"`
	}
)
