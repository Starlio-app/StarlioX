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
	app = fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			if code == fiber.StatusNotFound {
				return ctx.SendFile("./web/errors/404.html")
			}

			return nil
		},
	})

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

	api := app.Group("/api")

	update := api.Group("/update")
	get := api.Group("/get")
	create := api.Group("/create")

	update.Post("/settings", func(c *fiber.Ctx) error {
		return controllers.SettingsUpdate(c)
	})
	update.Post("/wallpaper", func(c *fiber.Ctx) error {
		return controllers.WallpaperUpdate(c)
	})
	update.Post("/startup", func(c *fiber.Ctx) error {
		return controllers.Startup(c)
	})

	create.Post("/label", func(c *fiber.Ctx) error {
		return controllers.CreateLabel(c)
	})

	get.Get("/settings", func(c *fiber.Ctx) error {
		return controllers.SettingsGet(c)
	})

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
