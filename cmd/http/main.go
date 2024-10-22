package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html/v2"
    "github.com/mehmetkmrc/kmrc_emlak/internal/infrastructure/adapters/web/handlers"
    "github.com/mehmetkmrc/kmrc_emlak/internal/infrastructure/adapters/web"
    "github.com/mehmetkmrc/kmrc_emlak/internal/infrastructure/configuration"
)

func main() {
    // Konfigürasyonu yükle
    config := configuration.LoadConfig()

    // HTML engine ayarı
    engine := html.New("./client/templates", ".html")

    // Fiber uygulamasını başlat
    app := fiber.New(fiber.Config{
        Views: engine,
    })

    // Static dosya yollarını tanımla
    web.SetupStaticFiles(app)

    // Route'ları setup et
    handlers.SetupRoutes(app)

    // Uygulamayı dinlemeye başla
    log.Fatal(app.Listen(config.ServerPort))
}
