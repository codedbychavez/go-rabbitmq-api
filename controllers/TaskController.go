package controllers

import (

	"github.com/codedbychavez/go-rabbitmq-api/models"

	"github.com/gofiber/fiber/v2"
)

type TaskController struct {

}

func NewTaskController() TaskController {
	return TaskController{}
}

func (ctrl TaskController) ReceiveTask(c *fiber.Ctx) error {
	params := new(struct {
		Title		string
		Description string
	})

	c.BodyParser(&params)

	if len(params.Title) == 0 || len(params.Description) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"ok":		false,
			"error": 	"Title or Decription not specified",
		})
	}

	task := models.Task{Title: params.Title, Description: params.Description}

	return c.JSON(fiber.Map{
		"ok": 	true,
		"task":	task,
	})
}