package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"ostkost/go-ps-hw-fiber/internal/users"
	"ostkost/go-ps-hw-fiber/pkg/tadapter"
	"ostkost/go-ps-hw-fiber/pkg/types"
	"ostkost/go-ps-hw-fiber/pkg/validator"
	"ostkost/go-ps-hw-fiber/views/components"

	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type ApiHandler struct {
	router       fiber.Router
	CustomLogger *slog.Logger
	userRepo     *users.UserRepository
}

func NewApiHandler(router fiber.Router, customLogger *slog.Logger, userRepo *users.UserRepository) {
	h := &ApiHandler{
		router:       router,
		CustomLogger: customLogger,
		userRepo:     userRepo,
	}
	apiGroup := h.router.Group("/api")
	apiGroup.Post("/register", h.register)
	apiGroup.Post("/login", h.login)
}

func (h ApiHandler) register(ctx *fiber.Ctx) error {
	f := types.RegisterForm{
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
		return tadapter.Render(ctx, component, http.StatusBadRequest)
	}
	// Db insert
	err := h.userRepo.Create(f)
	if err != nil {
		msg := fmt.Sprintf("Ошибка на сервере при создании пользователя")
		component := components.Notification(msg, components.NotificationError)
		return tadapter.Render(ctx, component, http.StatusInternalServerError)
	}
	msg := fmt.Sprintf("Регистрация успешна: %s", f.Email)
	component := components.Notification(msg, components.NotificationSuccess)
	return tadapter.Render(ctx, component, http.StatusCreated)
}

func (h ApiHandler) login(ctx *fiber.Ctx) error {
	f := types.LoginForm{
		Email:    ctx.FormValue("email"),
		Password: ctx.FormValue("password"),
	}
	errors := validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: f.Email, Message: "Неправильный email"},
		&validators.StringIsPresent{Name: "Password", Field: f.Password, Message: "Не заполнен пароль"},
	)
	if len(errors.Errors) > 0 {
		msg := validator.FormatErrors(errors)
		component := components.Notification(msg, components.NotificationError)
		return tadapter.Render(ctx, component, http.StatusBadRequest)
	}
	user := h.userRepo.GetByEmail(f.Email)
	if user == nil {
		return tadapter.Render(ctx, components.Notification("Неверный email или пароль", components.NotificationError), http.StatusUnauthorized)
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(f.Password))
	if err != nil {
		return tadapter.Render(ctx, components.Notification("Неверный email или пароль", components.NotificationError), http.StatusUnauthorized)
	}
	return ctx.SendStatus(http.StatusOK)
}
