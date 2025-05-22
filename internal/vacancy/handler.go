package vacancy

import (
	"ostkost/go-ps-fiber/pkg/tadatper"
	"ostkost/go-ps-fiber/pkg/validator"
	"ostkost/go-ps-fiber/views/components"
	"time"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type VacancyHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &VacancyHandler{
		router:       router,
		customLogger: customLogger,
	}
	vacancyGroup := h.router.Group("/vacancy")
	vacancyGroup.Post("/", h.vacancy)
}

func (h VacancyHandler) vacancy(ctx *fiber.Ctx) error {
	form := PostVacancyForm{
		Email:    ctx.FormValue("email"),
		Vacancy:  ctx.FormValue("vacancy"),
		Company:  ctx.FormValue("company"),
		Industry: ctx.FormValue("industry"),
		Salary:   ctx.FormValue("salary"),
		Location: ctx.FormValue("location"),
	}
	errors := validate.Validate(
		// Presence rules
		&validators.EmailIsPresent{Name: "Email", Field: form.Email, Message: "Не задан или неверный email"},
		&validators.StringIsPresent{Name: "Vacancy", Field: form.Vacancy, Message: "Не задана должность"},
		&validators.StringIsPresent{Name: "Company", Field: form.Company, Message: "Не задано название компании"},
		&validators.StringIsPresent{Name: "Industry", Field: form.Industry, Message: "Не задана сфера компаниии"},
		&validators.StringIsPresent{Name: "Salary", Field: form.Salary, Message: "Не задана заработная плата"},
		&validators.StringIsPresent{Name: "Location", Field: form.Location, Message: "Не задано расположение"},
		// Length rules
		&validators.StringLengthInRange{Name: "vacancy", Field: form.Vacancy, Min: 5, Max: 100, Message: "Название должности должно быть от 5 до 100 символов"},
		&validators.StringLengthInRange{Name: "company", Field: form.Company, Min: 2, Max: 100, Message: "Название компании должно быть от 2 до 100 символов"},
		&validators.StringLengthInRange{Name: "salary", Field: form.Salary, Min: 3, Max: 20, Message: "Зарплата должна быть от 3 до 20 символов"},
	)
	time.Sleep(time.Second * 2)
	var component templ.Component
	if len(errors.Errors) > 0 {
		msg := validator.FormatErrors(errors)
		component = components.Notification(msg, components.NotificationError)
		h.customLogger.Error().Msg("Post vacancy form errors")
		return tadatper.Render(ctx, component)
	}
	component = components.Notification("Вакансия создана", components.NotificationSuccess)
	return tadatper.Render(ctx, component)
}
