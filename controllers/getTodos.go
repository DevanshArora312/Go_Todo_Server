package controllers

import (
	"context"

	"github.com/devansharora312/go-server/config"
	// "github.com/devansharora312/go-server/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTodos(c *fiber.Ctx) error {
	var todos []Todo
	collection := config.Collection
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			return err
		}
		todos = append(todos, todo)
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Fetch Success!",
		"todos":   todos,
	})
}
