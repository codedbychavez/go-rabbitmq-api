package main

import (
	"github.com/codedbychavez/go-rabbitmq-api/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)


func main( ) {

	app := fiber.New()

	// Setup logger --> Logs requests to the console
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))


	taskController := controllers.NewTaskController()

	app.Post("/api/v1/sendtask", taskController.ReceiveTask)

	// default route for api
	app.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("Your GO API is working")
	})

	app.Listen(":3000")


}