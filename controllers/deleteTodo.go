package controllers

import (
	"context"

	"github.com/devansharora312/go-server/config"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTodo(c *fiber.Ctx) error {
	collection := config.Collection
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Invalid Todo ID!",
		})
	}

	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error in Deleting the Todo",
		})
	}

	return c.Status(202).JSON(fiber.Map{
		"message": "Delete Succesful!",
	})
}
