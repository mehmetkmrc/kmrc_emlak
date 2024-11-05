package converter

import (
	"github.com/mehmetkmrc/kmrc_emlak/_internal/core/domain/entity"
	"github.com/mehmetkmrc/kmrc_emlak/_internal/dto"
)

func GetUserModelToDto(userData *entity.User) *dto.GetUserResponse {
	return &dto.GetUserResponse{
		ID:        userData.ID,
		Name:      userData.Name,
		Surname:   userData.Surname,
		Email:     userData.Email,
		CreatedAt: userData.CreatedAt,
	}
}
