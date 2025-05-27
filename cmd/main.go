package main

import (
	"log"
	"ostkost/go-ps-fiber/internal/auth"
	"ostkost/go-ps-fiber/internal/home"
	"ostkost/go-ps-fiber/internal/vacancy"
	"ostkost/go-ps-fiber/pkg/config"
	"ostkost/go-ps-fiber/pkg/database"
	"ostkost/go-ps-fiber/pkg/logger"
	"ostkost/go-ps-fiber/pkg/middleware"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v3"
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
	// Store
	storage := postgres.New(postgres.Config{
		DB:         dbpool,
		Table:      "sessions",
		Reset:      false,
		GCInterval: 10 * time.Second,
	})
	store := session.New(session.Config{
		Storage: storage,
	})
	app.Use(middleware.AuthMiddleware(store))
	// Repositories
	vacancyRepo := vacancy.NewVacancyRepository(dbpool, customLogger)
	// Handlers
	home.NewHandler(app, customLogger, vacancyRepo, store)
	vacancy.NewHandler(app, customLogger, vacancyRepo)
	auth.NewAuthHandler(app, customLogger, store)
	// Init server
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
		return
	}
}
