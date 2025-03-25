package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"ostkost/go-ps-fiber/internal/home"
	"ostkost/go-ps-fiber/pkg/config"
)

func main() {
	config.Init()
	app := fiber.New()
	dbConfig := config.NewDatabaseConfig()
	log.Println(dbConfig)

	app.Use(recover.New())

	home.NewHandler(app)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
		return
	}
}
