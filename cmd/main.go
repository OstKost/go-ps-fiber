package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	slogfiber "github.com/samber/slog-fiber"
	"log"
	"ostkost/go-ps-hw-fiber/config"
	"ostkost/go-ps-hw-fiber/internal/logger"
	"ostkost/go-ps-hw-fiber/internal/pages"
	"strings"
)

func main() {
	config.Init()
	loggerConfig := config.NewLoggerConfig()
	dbConfig := config.NewDatabaseConfig()
	fmt.Println(dbConfig)

	engine := html.New("./html", ".html")
	engine.AddFuncMap(map[string]interface{}{
		"ToUpper": func(c string) string {
			return strings.ToUpper(c)
		},
	})

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	customLogger := logger.NewLogger(loggerConfig)
	app.Use(slogfiber.New(customLogger))

	app.Use(recover.New())

	pages.NewPagesHandler(app)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
