package converter

import (
	"github.com/gofrs/uuid"
	"github.com/mehmetkmrc/kmrc_emlak/internal/core/domain/entity"
	"github.com/mehmetkmrc/kmrc_emlak/internal/dto"
)

func GetSessionModelToDto(sessionData *entity.Sessions) *dto.SessionResponse {
	return &dto.SessionResponse{
		SessionId: uuid.UUID(sessionData.SessionId),
	}
}