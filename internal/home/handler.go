package home

import (
	"math"
	"net/http"
	"ostkost/go-ps-fiber/internal/vacancy"
	"ostkost/go-ps-fiber/pkg/tadatper"
	"ostkost/go-ps-fiber/views"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router            fiber.Router
	customLogger      *zerolog.Logger
	vacancyRepository *vacancy.VacancyRepository
	store             *session.Store
}

type User struct {
	Name string
	Age  int
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger, vacancyRepo *vacancy.VacancyRepository, store *session.Store) {
	h := &HomeHandler{
		router:            router,
		customLogger:      customLogger,
		vacancyRepository: vacancyRepo,
		store:             store,
	}
	h.router.Get("/", h.home)
	h.router.Get("/login", h.login)
	h.router.Get("/register", h.register)
	h.router.Get("/404", h.errorPage)
}

func (h HomeHandler) home(ctx *fiber.Ctx) error {
	PAGE_ITEMS := 2
	page := ctx.QueryInt("page")
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * PAGE_ITEMS
	count := h.vacancyRepository.CountAll()
	vacancies, err := h.vacancyRepository.GetAll(PAGE_ITEMS, offset)
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
		return ctx.SendStatus(http.StatusInternalServerError)
	}
	pagesCount := int(math.Ceil(float64(count / PAGE_ITEMS)))
	component := views.Main(vacancies, pagesCount, page)
	return tadatper.Render(ctx, component, http.StatusOK)
}

func (h HomeHandler) login(ctx *fiber.Ctx) error {
	component := views.Login()
	return tadatper.Render(ctx, component, http.StatusOK)
}

func (h HomeHandler) register(ctx *fiber.Ctx) error {
	component := views.Register()
	return tadatper.Render(ctx, component, http.StatusOK)
}

func (h HomeHandler) errorPage(ctx *fiber.Ctx) error {
	ctx.Status(http.StatusNotFound)
	return ctx.SendString("404")
}
