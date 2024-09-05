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

    log.Fatal(app.Listen(":3000"))
}