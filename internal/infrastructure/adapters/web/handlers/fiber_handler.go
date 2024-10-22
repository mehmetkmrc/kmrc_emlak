package handlers

import (
    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    app.Get("/", homePage)
    app.Get("/about", aboutPage)
    app.Get("/contacts", contactsPage)
    app.Get("/blogs", blogsPage)
    app.Get("/blog-single", blogSinglePage)
    app.Get("/projects", projectsPage)
}

func homePage(c *fiber.Ctx) error {
    return c.Render("home", fiber.Map{
        "Title": "Kömürcü Emlak - Anasayfa",
    })
}

func aboutPage(c *fiber.Ctx) error {
    return c.Render("about", fiber.Map{
        "Title": "Hakkımızda",
    })
}

func contactsPage(c *fiber.Ctx) error {
    return c.Render("contacts", fiber.Map{
        "Title": "İletişim",
    })
}

func blogsPage(c *fiber.Ctx) error {
    return c.Render("blogs", fiber.Map{
        "Title": "Haberler",
    })
}

func blogSinglePage(c *fiber.Ctx) error {
    return c.Render("blog-single", fiber.Map{
        "Title": "Tek Haber",
    })
}

func projectsPage(c *fiber.Ctx) error {
    return c.Render("projects", fiber.Map{
        "Title": "Projeler",
    })
}
