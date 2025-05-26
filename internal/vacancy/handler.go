package vacancy

import (
	"net/http"
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
	repository   *VacancyRepository
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger, reposotiry *VacancyRepository) {
	h := &VacancyHandler{
		router:       router,
		customLogger: customLogger,
		repository:   reposotiry,
	}
	vacancyGroup := h.router.Group("/vacancy")
	vacancyGroup.Post("/", h.createVacancy)
	vacancyGroup.Get("/", h.getAllVacancies)
}

func (h VacancyHandler) getAllVacancies(ctx *fiber.Ctx) error {
	PAGE_ITEMS := 2
	page := ctx.QueryInt("page")
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * PAGE_ITEMS
	vacancies, err := h.repository.GetAll(PAGE_ITEMS, offset)
	if err != nil {
		h.customLogger.Error().Err(err).Msg("DB: Failed to get vacancies")
		return err
	}
	return ctx.JSON(vacancies)
}

func (h VacancyHandler) createVacancy(ctx *fiber.Ctx) error {
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
	time.Sleep(time.Second * 2) // Simulate slow db
	var component templ.Component
	// Validation errors
	if len(errors.Errors) > 0 {
		msg := validator.FormatErrors(errors)
		component = components.Notification(msg, components.NotificationError)
		h.customLogger.Error().Msg("Validation: Post vacancy form errors")
		return tadatper.Render(ctx, component, http.StatusBadRequest)
	}
	// Db insert
	err := h.repository.Create(form)
	if err != nil {
		h.customLogger.Error().Err(err).Msg("DB: Failed to create vacancy")
		return tadatper.Render(ctx, components.Notification("Ошибка на сервере при создании вакансии", components.NotificationError), http.StatusBadRequest)
	}
	// Success response
	component = components.Notification("Вакансия создана", components.NotificationSuccess)
	return tadatper.Render(ctx, component, http.StatusOK)
}
