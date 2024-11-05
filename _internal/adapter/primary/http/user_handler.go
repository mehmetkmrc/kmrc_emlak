package http

import (
	"github.com/goccy/go-json"
	"github.com/mehmetkmrc/kmrc_emlak/_internal/converter"
	"github.com/mehmetkmrc/kmrc_emlak/_internal/dto"
	"time"

	"github.com/gofiber/fiber/v3"
)

func (s *server) Login(c fiber.Ctx) error {
	reqBody := new(dto.UserLoginRequest)
	body := c.Body()
	if err := json.Unmarshal(body, &reqBody); err != nil {
		return s.errorResponse(c, "error while trying to parse body", err, nil, fiber.StatusBadRequest)
	}

	userData, err := s.userService.Login(c.Context(), reqBody.Email, reqBody.Password)
	if err != nil {
		return s.errorResponse(c, "error while trying to login", err, nil, fiber.StatusBadRequest)
	}

	userResponse := converter.GetUserModelToDto(userData.User)
	bearerAccess := "Bearer " + userData.AccessToken
	c.Cookie(&fiber.Cookie{
		Name:     "id",
		Value:    userData.User.ID,
		Expires:  time.Now().Add(3 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
	})
	c.Cookie(&fiber.Cookie{
		Name:     "name",
		Value:    userData.User.Name + " " + userData.User.Surname,
		Expires:  time.Now().Add(3 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
	})
	c.Cookie(&fiber.Cookie{
		Name:     AccessToken,
		Value:    bearerAccess,
		Expires:  time.Now().Add(time.Hour * 3),
		HTTPOnly: true,
		Secure:   true,
	})

	c.Cookie(&fiber.Cookie{
		Name:     AccessPublic,
		Value:    userData.AccessPublic,
		Expires:  time.Now().Add(time.Hour * 3),
		HTTPOnly: true,
		Secure:   true,
	})

	bearerRefresh := "Bearer " + userData.RefreshToken
	c.Cookie(&fiber.Cookie{
		Name:     RefreshToken,
		Value:    bearerRefresh,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
	})

	c.Cookie(&fiber.Cookie{
		Name:     RefreshPublic,
		Value:    userData.RefreshPublic,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
	})

	return s.successResponse(c, userResponse, "user logged in successfully", fiber.StatusOK)
}
