package home

import (
	"github.com/gofiber/fiber/v2"
	"ostkost/go-ps-fiber/pkg/tadatper"
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
	component := views.Main()
	return tadatper.Render(ctx, component)
}
