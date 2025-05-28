package robots

import "github.com/gofiber/fiber/v2"

type RobotsHandler struct {
	router fiber.Router
}

func NewRobotsHandler(router fiber.Router) {
	h := &RobotsHandler{
		router: router,
	}
	h.router.Get("/robots.txt", h.robots)
}

func (h RobotsHandler) robots(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "text/plain")
	return ctx.SendFile("./public/robots.txt")
}
