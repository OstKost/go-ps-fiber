package home

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"ostkost/go-ps-fiber/views"
)

type HomeHandler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	h := &HomeHandler{
		router: router,
	}
	h.router.Get("/", h.home)
}

type User struct {
	Name string
	Age  int
}

func (h HomeHandler) home(ctx *fiber.Ctx) error {
	component := views.Hello("MEDIA")
	err := component.Render(context.Background(), os.Stdout)
	if err != nil {
		log.Println(err)
		return ctx.SendString(err.Error())
	}
	return ctx.SendString("Hello World")
}
