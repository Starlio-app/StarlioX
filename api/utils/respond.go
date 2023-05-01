package utils

import (
	"encoding/json"
	"github.com/Redume/EveryNASA/utils"

	"github.com/gofiber/fiber/v2"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(c *fiber.Ctx, data map[string]interface{}) {
	c.Set("Content-Type", "application/json")
	err := json.NewEncoder(c).Encode(data)
	if err != nil {
		utils.Log(err.Error())
	}
}
