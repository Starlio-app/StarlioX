package main

import (
	"github.com/Redume/EveryNasa/api/controllers"
	"github.com/Redume/EveryNasa/utils"
	"github.com/getlantern/systray"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	go open.Run("http://localhost:3000")
	go utils.Database()
	go systray.Run(utils.Tray, utils.Quit)
	go utils.StartWallpaper()

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
		con := utils.Connected()
		if con == false {
			return c.SendFile("./web/errors/500.html")
		}

		return c.SendFile("./web/page/gallery.html")
	})
	app.Get("/settings", func(c *fiber.Ctx) error {
		con := utils.Connected()
		if con == false {
			return c.SendFile("./web/errors/500.html")
		}

		return c.SendFile("./web/page/settings.html")
	})
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendFile("./web/page/about.html")
	})
	app.Get("/favorite", func(c *fiber.Ctx) error {
		return c.SendFile("./web/page/favorite.html")
	})

	api := app.Group("/api")

	update := api.Group("/update")
	get := api.Group("/get")
	add := api.Group("/add")
	del := api.Group("/del")
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

	get.Get("/favorites", func(c *fiber.Ctx) error {
		return controllers.GetFavorites(c)
	})

	add.Post("/favorite", func(c *fiber.Ctx) error {
		return controllers.AddFavorite(c)
	})

	del.Post("/favorite", func(c *fiber.Ctx) error {
		return controllers.DeleteFavorite(c)
	})

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
