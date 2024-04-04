package main

import (
	"fmt"
	"golang/database"
	"golang/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello Naveen")
}

func initDatabase() error {
	dsn := "host=localhost user=postgres password=root dbname=p1 port=5432 sslmode=disable"

	var err error
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Println("Database connected!")
	return database.DBConn.AutoMigrate(&models.Todo{})
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	app.Get("/todos", models.GetTodos)
	app.Post("/todos", models.CreateTodo)
	app.Get("/todos/:id", models.GetTodoById)
	app.Put("/todos/:id", models.UpdateTodo)
	app.Delete("/todos/:id", models.DeleteTodo)
}

func main() {
	app := fiber.New()

	app.Use(cors.New())

	err := initDatabase()
	if err != nil {
		panic("Failed to connect to database!")
	}

	setupRoutes(app)

	err = app.Listen(":8000")
	if err != nil {
		panic(err)
	}
}
