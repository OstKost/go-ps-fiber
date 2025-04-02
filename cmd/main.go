package main

import (
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"ostkost/go-ps-fiber/internal/home"
	"ostkost/go-ps-fiber/pkg/config"
	"ostkost/go-ps-fiber/pkg/logger"
)

func main() {
	config.Init()
	app := fiber.New()
	logConfig := config.NewLogConfig()
	dbConfig := config.NewDatabaseConfig()
	customLogger := logger.NewLogger(logConfig)
	log.Println(dbConfig)

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New())

	home.NewHandler(app)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
		return
	}
}
