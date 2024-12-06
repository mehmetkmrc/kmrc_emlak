package converter

import (
	"github.com/mehmetkmrc/kmrc_emlak/internal/core/domain/entity"
	"github.com/mehmetkmrc/kmrc_emlak/internal/dto"
)

func GetUserModelToDto(userData *entity.User) *dto.GetUserResponse {
	return &dto.GetUserResponse{
		UserID:    userData.UserID,
		Name:      userData.Name,
		Surname:   userData.Surname,
		Email:     userData.Email,
		CreatedAt: userData.CreatedAt,
	}
}
