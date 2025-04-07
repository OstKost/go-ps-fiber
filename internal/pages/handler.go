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
	h.router.Get("/page", h.about)
}

func (h PagesHandler) base(ctx *fiber.Ctx) error {
	slog.Info("base", "page", "base", "count", 1)
	return ctx.JSON(fiber.Map{"status": "OK"})
}

func (h PagesHandler) about(ctx *fiber.Ctx) error {
	links := []struct {
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
		Title string
		Links []struct {
			URL  string
			Name string
		}
	}{
		Title: "Главная страница",
		Links: links,
	}
	return ctx.Render("page", data)
}
