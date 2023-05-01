package main

import (
	"github.com/Redume/EveryNASA/api/controllers"
	"github.com/Redume/EveryNASA/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
)

func main() {
	go utils.CheckLogs()
	app := fiber.New(fiber.Config{
		AppName: "EveryNASA",
	})
	app.Static("/static", "./interface/static")
	app.Use(favicon.New(favicon.Config{
		File: "./interface/static/assets/icons/favicon.ico",
		URL:  "/favicon.ico",
	}))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("./interface/page/gallery.html")
	})

	app.Get("/about", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("./interface/page/about.html")
	})

	api := app.Group("/api")
	update := api.Group("/update")

	update.Post("/wallpaper", func(ctx *fiber.Ctx) error {
		return controllers.SetWallpaper(ctx)
	})

	err := app.Listen(":4000")
	if err != nil {
		utils.Log(err.Error())
	}
}
