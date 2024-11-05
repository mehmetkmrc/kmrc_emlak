package http

import (
	"github.com/gofiber/fiber/v3"
)

func (s *server) LoginWeb(c fiber.Ctx) error {
	path := "login"
	return c.Render(path, fiber.Map{
		"Title": "Login",
	})
}

