package main

import (
	"fmt"
	"log"
	"ostkost/go-ps-fiber/internal/home"
	"ostkost/go-ps-fiber/internal/vacancy"
	"ostkost/go-ps-fiber/pkg/config"
	"ostkost/go-ps-fiber/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	logConfig := config.NewLogConfig()
	dbConfig := config.NewDatabaseConfig()
	customLogger := logger.NewLogger(logConfig)
	// engine := html.New("./html", ".html")

	fmt.Println("dbConfig:", dbConfig)

	// app := fiber.New(fiber.Config{
	// 	Views: engine,
	// })
	app := fiber.New()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New())

	app.Static("/public", "./public")

	home.NewHandler(app)
	vacancy.NewHandler(app, customLogger)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
		return
	}
}
