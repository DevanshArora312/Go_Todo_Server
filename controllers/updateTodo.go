package controllers

import (
	"context"

	"github.com/devansharora312/go-server/config"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateTodo(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Invalid Todo ID!",
		})
	}
	collection := config.Collection

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"completed": true}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error in updating the Todo",
		})
	}
	// defer cursor.Close(context.Background())
	return c.Status(202).JSON(fiber.Map{
		"message": "Update Succesful!",
	})
}
