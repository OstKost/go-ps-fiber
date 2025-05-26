package main

import (
	"log"
	"ostkost/go-ps-fiber/internal/home"
	"ostkost/go-ps-fiber/internal/vacancy"
	"ostkost/go-ps-fiber/pkg/config"
	"ostkost/go-ps-fiber/pkg/database"
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
	// Init Fiber App
	app := fiber.New()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New())
	// Static
	app.Static("/public", "./public")
	// Database
	dbpool := database.CreateDbPool(dbConfig, customLogger)
	defer dbpool.Close()
	// Repositories
	vacancyRepo := vacancy.NewVacancyRepository(dbpool, customLogger)
	// Handlers
	home.NewHandler(app, customLogger, vacancyRepo)
	vacancy.NewHandler(app, customLogger, vacancyRepo)
	// Init server
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
		return
	}
}
