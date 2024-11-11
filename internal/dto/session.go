package dto

import "github.com/gofrs/uuid"

type (
	SessionResponse struct {
		SessionId uuid.UUID `json:"session_id"`
	}
)