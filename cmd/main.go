package main

import (
	"fmt"
	"log"
	"ostkost/go-ps-hw-fiber/config"
	"ostkost/go-ps-hw-fiber/internal/api"
	"ostkost/go-ps-hw-fiber/internal/pages"
	"ostkost/go-ps-hw-fiber/internal/users"
	"ostkost/go-ps-hw-fiber/pkg/database"
	"ostkost/go-ps-hw-fiber/pkg/logger"
	"ostkost/go-ps-hw-fiber/pkg/middleware"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v3"
	sf "github.com/samber/slog-fiber"
)

func main() {
	config.Init()
	loggerConfig := config.NewLoggerConfig()
	dbConfig := config.NewDatabaseConfig()
	fmt.Println(dbConfig)

	app := fiber.New()

	customLogger := logger.NewLogger(loggerConfig)
	app.Use(sf.New(customLogger))
	// Recover
	app.Use(recover.New())
	// Static
	app.Static("/public", "./public")
	// Database
	dbpool := database.CreateDbPool(dbConfig)
	storage := postgres.New(postgres.Config{
		DB:         dbpool,
		Table:      "news.sessions",
		Reset:      false,
		GCInterval: 10 * time.Second,
	})
	sessionStore := session.New(session.Config{
		Storage: storage,
	})
	app.Use(middleware.AuthMiddleware(sessionStore))
	// Repositories
	userRepo := users.NewUserRepository(dbpool, customLogger)
	// Handlers
	pages.NewPagesHandler(app)
	api.NewApiHandler(app, customLogger, userRepo, sessionStore)
	users.NewUserHandler(app, userRepo)
	// Init server
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
