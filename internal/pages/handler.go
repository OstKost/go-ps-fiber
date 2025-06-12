package pages

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type PagesHandler struct {
	router fiber.Router
}

func NewPagesHandler(router fiber.Router) {
	h := &PagesHandler{
		router: router,
	}

	h.router.Get("/", h.base)
	h.router.Get("/page", h.about)
}

func (h PagesHandler) base(ctx *fiber.Ctx) error {
	slog.Info("base", "page", "base", "count", 1)
	return ctx.JSON(fiber.Map{"status": "OK"})
}

func (h PagesHandler) about(ctx *fiber.Ctx) error {
	categories := []struct {
		URL  string
		Name string
	}{
		{URL: "/", Name: "base"},
		{URL: "/page", Name: "page"},
		{URL: "/page/users", Name: "users"},
		{URL: "/page/users/:id", Name: "user"},
		{URL: "/page/groups", Name: "groups"},
	}
	data := struct {
		Title      string
		Categories []struct {
			URL  string
			Name string
		}
	}{
		Title:      "Главная страница",
		Categories: categories,
	}
	return ctx.Render("page", data)
}
