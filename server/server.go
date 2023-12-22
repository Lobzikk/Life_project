package server

import (
	"github.com/gofiber/fiber/v2"
	"life_project/json_resolve"
	"life_project/life"
	"log"
)

func SendMap(ctx *fiber.Ctx, world life.World) error {
	res, err := jsonresolve.MapToJSON(world)
	if err != nil {
		return err
	}
	err = ctx.Send(res)
	return err
}

func StopServer(app *fiber.App) error {
	err := app.Server().Shutdown()
	log.Println("Server is shutting down!")
	return err
}

func Listen(app *fiber.App, world life.World) error {
	app.Get("/getworld", func(c *fiber.Ctx) error {
		return SendMap(c, world)
	})
	app.Get("/stop", func(c *fiber.Ctx) error {
		return StopServer(app)
	})
	return app.Listen(":8080")
}
