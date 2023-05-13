package main

import (
	"github.com/10wpressure/simple-rest-api/database"
	"github.com/10wpressure/simple-rest-api/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	database.ConnectDb()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	setupRoutes(app)

	app.Static("/", "./public")

	app.Use(handlers.NotFound)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
