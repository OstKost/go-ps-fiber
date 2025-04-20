package pages

import (
	"ostkost/go-ps-hw-fiber/pkg/tadapter"
	"ostkost/go-ps-hw-fiber/views"
	"ostkost/go-ps-hw-fiber/views/components"

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
}

func (h PagesHandler) index(ctx *fiber.Ctx) error {
	component := views.Index()
	return tadapter.Render(ctx, component)
}

func (h PagesHandler) categories(ctx *fiber.Ctx) error {

	links := []components.NavbarItemProps{
		{Href: "/", Text: "Eда", Img: "food.jpeg"},
		{Href: "/", Text: "Животные", Img: "animals.jpeg"},
		{Href: "/", Text: "Машины", Img: "cars.jpeg"},
		{Href: "/", Text: "Спорт", Img: "sport.jpeg"},
		{Href: "/", Text: "Музыка", Img: "music.jpeg"},
		{Href: "/", Text: "Технологии", Img: "technology.jpeg"},
		{Href: "/", Text: "Прочее", Img: "other.jpeg"},
	}

	component := views.Categories(views.CategoriesProps{NavItems: links})
	return tadapter.Render(ctx, component)
}
