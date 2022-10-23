package page

import (
	"github.com/Redume/EveryNasa/functions"
	"github.com/gofiber/fiber/v2"
)

func Gallery(c *fiber.Ctx) error {
	con := functions.Connected()
	if con == false {
		return c.SendFile("./web/errors/500.html")
	}

	return c.SendFile("./web/src/gallery.html")
}
