package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
	"log"
	"ostkost/go-ps-hw-fiber/config"
	"ostkost/go-ps-hw-fiber/internal/logger"
	"ostkost/go-ps-hw-fiber/internal/pages"
)

func main() {
	config.Init()
	loggerConfig := config.NewLoggerConfig()
	dbConfig := config.NewDatabaseConfig()
	fmt.Println(dbConfig)

	customLogger := logger.NewLogger(loggerConfig)

	app := fiber.New()

	app.Use(slogfiber.New(customLogger))
	app.Use(recover.New())

	pages.NewPagesHandler(app)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
