package home

import "github.com/gofiber/fiber/v2"

type HomeHandler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	h := &HomeHandler{
		router: router,
	}
	h.router.Get("/", h.home)
}

func (h HomeHandler) home(ctx *fiber.Ctx) error {
	ctx.Hostname()
	return ctx.SendString("OK")
}
