package http

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

func (s *server) SetupRouter() {
	s.webSetUp()
	s.authSetUp()
}

func (s *server) webSetUp() {
	s.app.Get("/", s.HomeWeb)
	// s.app.Get("/", func(c fiber.Ctx) error{
	// 	return c.Redirect().To("/login")
	// })
	s.app.Get("/ping", func(c fiber.Ctx) error{
		return c.SendString("Pong")
	})
	s.app.Get("/login", s.LoginWeb, s.RateLimiter(5, time.Minute))

}

func (s *server) authSetUp() {
	route := s.app.Group("/auth")
	route.Post("/login", s.Login, s.RateLimiter(5, time.Minute), s.LoginValidation)
}
