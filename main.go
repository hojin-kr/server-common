package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hojin-kr/common"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/common/:uuid", func(c *fiber.Ctx) error {
		common := common.Common{}
		common.UUID = c.Params("uuid")
		common.GetCommon()
		return c.JSON(common)
	})

	// update
	app.Post("/common/:uuid", func(c *fiber.Ctx) error {
		common := common.Common{}
		common.UUID = c.Params("uuid")
		common.GetCommon()
		if common.UUID == "" {
			return c.JSON(common)
		}
		// c.Body() to common
		c.BodyParser(&common)
		common.SetCommon()

		return c.JSON(common)

	})

	app.Listen(":3000")
}
