package main

import (
	"github.com/codedbychavez/go-rabbitmq-api/controllers"
	"github.com/gofiber/fiber/v2"
)


func main( ) {
	app := fiber.New()
	taskController := controllers.NewTaskController()

	app.Post("/api/v1/sendtask", taskController.ReceiveTask)

	// default route for api
	app.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("Your GO API is working")
	})

	app.Listen(":3000")
}