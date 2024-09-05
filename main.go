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
    app.Get("/", func(c *fiber.Ctx) error {
        return c.Render("home", fiber.Map{
            "Title": "Go Fiber Template Example",
        });
    })

    log.Fatal(app.Listen(":3000"))
}