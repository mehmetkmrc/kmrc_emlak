package entity

import (
	"time"

	"github.com/google/uuid"
)

type (
	Sessions struct {
		SessionId uuid.UUID `json:"session_id"`
		UserId	  uuid.UUID `json:"user_id"`
		Token 	  string	`json:"token"`
		IPAdress string		`json:"ip_address"`
		UserAgent string 	`json:"user_agent"`
		CreatedAt time.Time `json:"created_at"`
		ExpiresAt time.Time `json:"expires_at"`
		LastAccess time.Time `json:"last_access"`
		IsActive bool 		 `json:"is_active"`
		Location string 	 `json:"location"`
	}
)