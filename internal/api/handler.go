package api

import (
	"ostkost/go-ps-hw-fiber/pkg/tadapter"
	"ostkost/go-ps-hw-fiber/pkg/validator"
	"ostkost/go-ps-hw-fiber/views/components"

	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
)

type ApiHandler struct {
	router fiber.Router
}

func NewApiHandler(router fiber.Router) {
	h := &ApiHandler{
		router: router,
	}

	apiGroup := h.router.Group("/api")
	apiGroup.Post("/register", h.register)
}

func (h ApiHandler) register(ctx *fiber.Ctx) error {
	f := RegisterForm{
		Name:     ctx.FormValue("name"),
		Email:    ctx.FormValue("email"),
		Password: ctx.FormValue("password"),
	}

	errors := validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: f.Email, Message: "Неправильный email"},
		&validators.StringIsPresent{Name: "Password", Field: f.Password, Message: "Не заполнен пароль"},
		&validators.StringIsPresent{Name: "Name", Field: f.Name, Message: "Не заполнено имя"},
	)

	if len(errors.Errors) > 0 {
		msg := validator.FormatErrors(errors)
		component := components.Notification(msg, components.NotificationError)
		return tadapter.Render(ctx, component)

	}

	component := components.Notification("Регистрация успешна", components.NotificationSuccess)
	return tadapter.Render(ctx, component)
}
