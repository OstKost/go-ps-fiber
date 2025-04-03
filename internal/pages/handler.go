package pages

import (
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

type PagesHandler struct {
	router fiber.Router
}

func NewPagesHandler(router fiber.Router) {
	h := &PagesHandler{
		router: router,
	}

	h.router.Get("/", h.base)
	h.router.Get("/about", h.about)
}

func (h PagesHandler) base(ctx *fiber.Ctx) error {
	slog.Info("base", "page", "base", "count", 1)
	return ctx.JSON(fiber.Map{"status": "OK"})
}

func (h PagesHandler) about(ctx *fiber.Ctx) error {
	slog.Error("about", "page", "about", "count", 2)
	return ctx.JSON(fiber.Map{"page": "about", "status": "OK"})
}
