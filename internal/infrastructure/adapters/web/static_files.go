package web

import "github.com/gofiber/fiber/v2"

func SetupStaticFiles(app *fiber.App) {
    app.Static("/client/public/css", "../../client/public/css")
    app.Static("/client/public/js", "../../client/public/js")
    app.Static("/client/public/fonts", "../../client/public/fonts")
    app.Static("/client/public/images", "../../client/public/images")
    app.Static("client/public/videos", "../../client/public/videos")
}
