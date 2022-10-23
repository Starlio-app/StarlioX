package main

import (
	"github.com/Redume/EveryNasa/api/controllers"
	"github.com/Redume/EveryNasa/functions"
	"github.com/Redume/EveryNasa/web/page"

	"github.com/getlantern/systray"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	go functions.Database()
	go systray.Run(functions.Tray, functions.Quit)
	go functions.StartWallpaper()

	app := fiber.New()
	app.Static("/static", "./web/static")
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return page.Gallery(c)
	})
	app.Get("/settings", func(c *fiber.Ctx) error {
		return page.Settings(c)
	})
	app.Get("/about", func(c *fiber.Ctx) error {
		return page.About(c)
	})

	app.Post("/api/update/settings", func(c *fiber.Ctx) error {
		return controllers.SettingsUpdate(c)
	})
	app.Post("/api/update/wallpaper", func(c *fiber.Ctx) error {
		return controllers.WallpaperUpdate(c)
	})
	app.Post("/api/update/startup", func(c *fiber.Ctx) error {
		return controllers.Startup(c)
	})
	app.Post("/api/create/label", func(c *fiber.Ctx) error {
		return controllers.CreateLabel(c)
	})

	app.Get("/api/get/settings", func(c *fiber.Ctx) error {
		return controllers.SettingsGet(c)
	})

	app.Use(func(c *fiber.Ctx) error {
		err := c.SendStatus(404)
		if err != nil {
			functions.Logger(err.Error())
		}
		return c.SendFile("./web/errors/404.html")
	})

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
