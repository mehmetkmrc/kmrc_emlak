package http

import (
	"github.com/goccy/go-json"
	"github.com/mehmetkmrc/kmrc_emlak/_internal/dto"

	"github.com/gofiber/fiber/v3"
)

func (s *server) LoginValidation(c fiber.Ctx) error {
	data := new(dto.UserLoginRequest)
	body := c.Body()
	err := json.Unmarshal(body, &data)
	if err != nil {
		return s.errorResponse(c, "invalid request body", err, nil, fiber.StatusBadRequest)
	}

	validationErrors := ValidateRequestByStruct(data)
	if len(validationErrors) > 0 {
		return s.errorResponse(c, "validation failed", nil, validationErrors, fiber.StatusUnprocessableEntity)
	}

	return c.Next()
}
