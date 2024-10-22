package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/mehmetkmrc/kmrc_emlak/internal/infrastructure/adapters/web"
	"github.com/mehmetkmrc/kmrc_emlak/internal/infrastructure/adapters/web/handlers"
	"github.com/mehmetkmrc/kmrc_emlak/internal/infrastructure/configuration"
)

func main() {
    // Konfigürasyonu yükle
    config := configuration.LoadConfig()

    // Template engine'i tanımla ve template dosyalarının yolunu belirt
    engine := html.New("../../client/templates", ".html")

    // Fiber uygulamasını başlat
    app := fiber.New(fiber.Config{
        Views: engine,
    })

    // Static dosya ayarları
    web.SetupStaticFiles(app)

    // Routerları ayarla
    handlers.SetupRoutes(app)

    // Uygulamayı çalıştır
    log.Fatal(app.Listen(config.ServerPort))
}
