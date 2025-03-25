package pages

import "github.com/gofiber/fiber/v2"

type PagesHandler struct {
	router fiber.Router
}

func NewPagesHandler(router fiber.Router) {
	h := &PagesHandler{
		router: router,
	}

	h.router.Get("/", h.base)
}

func (h PagesHandler) base(ctx *fiber.Ctx) error {
	return ctx.JSON("OK")
}
