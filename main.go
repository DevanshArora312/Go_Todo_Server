package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/devansharora312/go-server/config"
	"github.com/devansharora312/go-server/controllers"
	"github.com/devansharora312/go-server/models"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

var allTodos = []models.Todo{}

func main() {
	app := fiber.New()

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	config.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	app.Get("/api/todos", controllers.GetTodos)
	app.Get("/api/todos/:id", controllers.GetOneBrute)
	app.Post("/api/create", controllers.CreateTodo)
	app.Patch("/api/update/:id", controllers.UpdateTodo)
	app.Delete("/api/delete/:id", controllers.DeleteTodo)
	defer config.Client.Disconnect(context.Background())

	log.Fatal(app.Listen(fmt.Sprintf(":%v", PORT)))
}
