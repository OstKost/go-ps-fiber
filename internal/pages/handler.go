package pages

import (
	"ostkost/go-ps-hw-fiber/pkg/tadapter"

	"github.com/gofiber/fiber/v2"
)

type PagesHandler struct {
	router fiber.Router
}

func NewPagesHandler(router fiber.Router) {
	h := &PagesHandler{
		router: router,
	}

	h.router.Get("/", h.index)
	h.router.Get("/categories", h.categories)
	h.router.Get("/register", h.register)
}

func (h PagesHandler) index(ctx *fiber.Ctx) error {
	component := IndexComponent()
	return tadapter.Render(ctx, component)
}

func (h PagesHandler) categories(ctx *fiber.Ctx) error {
	component := CategoriesComponent()
	return tadapter.Render(ctx, component)
}

func (h PagesHandler) register(ctx *fiber.Ctx) error {
	component := RegisterComponent()
	return tadapter.Render(ctx, component)
}
