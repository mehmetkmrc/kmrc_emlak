package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main(){
    
    
    engine := html.New("./client/templates", ".html")


    app := fiber.New(fiber.Config{
        Views: engine,
    })

    app.Static("/client/public/css", "./client/public/css")
    app.Static("/client/public/js", "./client/public/js")
    app.Static("/client/public/fonts", "./client/public/fonts")
    
    app.Get("/", func(c *fiber.Ctx) error {
        return c.Render("home", fiber.Map{
            "Title": "Kömürcü Emlak - Anasayfa",
        });
    })
    app.Get("/about", func (c *fiber.Ctx) error {
        return c.Render("about", fiber.Map{
            "Title": "Hakkımızda",
        })
    })
    app.Get("/contacts", func (c *fiber.Ctx) error {
        return c.Render("contacts", fiber.Map{
            "Title": "İletişim",
        })
    })
    app.Get("/blogs", func (c *fiber.Ctx) error{
        return c.Render("blogs", fiber.Map{
            "Title": "Haberler",
        })
    })
    app.Get("/blog-single", func (c *fiber.Ctx) error {
        return c.Render("blog-single", fiber.Map{
            "Title": "Tek Haberler",
        })
    })
    app.Get("/projects", func (c *fiber.Ctx) error {
        return c.Render("projects", fiber.Map{
            "Title": "Projeler",
        })
    })

    log.Fatal(app.Listen(":3000"))
}