package home

import (
	"ostkost/go-ps-fiber/pkg/tadatper"
	"ostkost/go-ps-fiber/views"

	"github.com/gofiber/fiber/v2"
)

type HomeHandler struct {
	router fiber.Router
}

type User struct {
	Name string
	Age  int
}

func NewHandler(router fiber.Router) {
	h := &HomeHandler{
		router: router,
	}
	h.router.Get("/", h.home)
}

func (h HomeHandler) home(ctx *fiber.Ctx) error {
	component := views.Main()
	return tadatper.Render(ctx, component)
}
