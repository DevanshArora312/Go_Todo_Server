package controllers

import (
	"context"

	"github.com/devansharora312/go-server/config"
	// "github.com/devansharora312/go-server/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOneBrute(c *fiber.Ctx) error {
	collection := config.Collection
	id, err1 := primitive.ObjectIDFromHex(c.Params("id"))
	if err1 != nil {
		return err1
	}
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
		// fmt.Printf("%v \n",todo.ID)
		if todo.ID == id {
			return c.Status(200).JSON(fiber.Map{
				"message": "Fetch Success!",
				"todo":    todo,
			})
		}
	}
	return c.Status(404).JSON(fiber.Map{
		"message": "No todo with given ID found!",
	})
}

func GetOne(c *fiber.Ctx) error {
	collection := config.Collection
	id, err1 := primitive.ObjectIDFromHex(c.Params("id"))
	if err1 != nil {
		return err1
	}
	todo := collection.FindOne(context.Background(), bson.M{"_id": id})

	if todo == nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "No todo with given ID found!",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Fetch Success!",
		"todo":    todo,
	})

}
