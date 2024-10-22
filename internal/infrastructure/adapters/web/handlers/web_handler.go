package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    app.Get("/", HomeHandler)
    app.Get("/about", AboutHandler)
    app.Get("/contacts", ContactsHandler)
    app.Get("/blogs", BlogsHandler)
    app.Get("/blog-single", BlogSingleHandler)
    app.Get("/projects", ProjectsHandler)
}

func HomeHandler(c *fiber.Ctx) error {
    return c.Render("home", fiber.Map{
        "Title": "Kömürcü Emlak - Anasayfa",
    })
}

func AboutHandler(c *fiber.Ctx) error {
    return c.Render("about", fiber.Map{
        "Title": "Hakkımızda",
    })
}

func ContactsHandler(c *fiber.Ctx) error {
    return c.Render("contacts", fiber.Map{
        "Title": "İletişim",
    })
}

func BlogsHandler(c *fiber.Ctx) error {
    return c.Render("blogs", fiber.Map{
        "Title": "Haberler",
    })
}

func BlogSingleHandler(c *fiber.Ctx) error {
    return c.Render("blog-single", fiber.Map{
        "Title": "Tek Haberler",
    })
}

func ProjectsHandler(c *fiber.Ctx) error {
    return c.Render("projects", fiber.Map{
        "Title": "Projeler",
    })
}
