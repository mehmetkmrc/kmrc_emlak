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
		PhoneNumber string    `json:"phone_number"`
		CreatedAt   time.Time `json:"created_at"`
	}
)
