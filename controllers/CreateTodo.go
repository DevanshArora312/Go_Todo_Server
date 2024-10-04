package controllers

import (
	"context"

	"github.com/devansharora312/go-server/config"
	"github.com/devansharora312/go-server/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo = models.Todo

func CreateTodo(c *fiber.Ctx) error {
	todo := &Todo{}
	if err := c.BodyParser(todo); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error trying to read request body",
		})
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Todo Body can't be empty",
		})
	}
	collection := config.Collection
	insertResult, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error trying to insert Tofo to DB",
		})
	}
	todo.ID = insertResult.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(fiber.Map{
		"message": "Added Todo",
		"todos":   todo,
	})
}
