package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"ostkost/go-ps-hw-fiber/config"
	"ostkost/go-ps-hw-fiber/internal/pages"
)

func main() {
	config.Init()
	dbConfig := config.NewDatabaseConfig()

	fmt.Println(dbConfig)

	app := fiber.New()
	app.Use(recover.New())

	pages.NewPagesHandler(app)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
