package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	ID        int    `json:"_id" bson:"_id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

var allTodos = []Todo{}

var collection *mongo.Collection

func main() {
	app := fiber.New()

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	// MONGO_URI := os.Getenv("MONGO_URI")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	app.Get("/api/todos", getAll)
	app.Get("/api/todos/:id", getOne)
	app.Post("/api/create", CreateTodo)

	log.Fatal(app.Listen(fmt.Sprintf(":%v", PORT)))
}

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
	todo.ID = len(allTodos) + 1
	allTodos = append(allTodos, *todo)
	return c.Status(201).JSON(fiber.Map{
		"message": "Added Todo",
		"todos":   allTodos,
	})
}

func getAll(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"todos": allTodos,
	})
}

func getOne(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	todo := Todo{}
	for _, item := range allTodos {
		if item.ID == id {
			todo = item
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
