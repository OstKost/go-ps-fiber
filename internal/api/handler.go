package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"ostkost/go-ps-hw-fiber/internal/users"
	sess "ostkost/go-ps-hw-fiber/pkg/session"
	"ostkost/go-ps-hw-fiber/pkg/tadapter"
	"ostkost/go-ps-hw-fiber/pkg/types"
	"ostkost/go-ps-hw-fiber/pkg/validator"
	"ostkost/go-ps-hw-fiber/views/components"

	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	s "github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
)

type ApiHandler struct {
	router       fiber.Router
	CustomLogger *slog.Logger
	userRepo     *users.UserRepository
	sessionStore *s.Store
}

func NewApiHandler(router fiber.Router, customLogger *slog.Logger, userRepo *users.UserRepository, session *s.Store) {
	h := &ApiHandler{
		router:       router,
		CustomLogger: customLogger,
		userRepo:     userRepo,
		sessionStore: session,
	}
	apiGroup := h.router.Group("/api")
	apiGroup.Post("/register", h.register)
	apiGroup.Post("/login", h.login)
	apiGroup.Post("/logout", h.logout)
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
		&validators.StringLengthInRange{Name: "name", Field: f.Name, Min: 2, Max: 50, Message: "Имя должно быть от 2 до 50 символов"},
		&validators.StringLengthInRange{Name: "password", Field: f.Password, Min: 6, Max: 20, Message: "Пароль должен быть от 6 до 20 символов"},
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
	user := h.userRepo.GetByEmail(f.Email)
	if user == nil {
		return tadapter.Render(ctx, components.Notification("Ошибка получения пользователя", components.NotificationError), http.StatusInternalServerError)
	}
	// Session
	session, err := sess.GetSession(ctx, h.sessionStore, h.CustomLogger)
	if err != nil {
		return tadapter.Render(ctx, components.Notification("Ошибка получения сессии", components.NotificationError), http.StatusInternalServerError)
	}
	session.Set("email", f.Email)
	session.Set("userId", user.Id)
	session.Set("name", user.Name)
	err = sess.SaveSession(session, h.CustomLogger)
	if err != nil {
		return tadapter.Render(ctx, components.Notification("Ошибка сохранения сессии", components.NotificationError), http.StatusInternalServerError)
	}
	ctx.Response().Header.Add("Hx-Redirect", "/")
	return ctx.Redirect("/", http.StatusOK)
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
	session, err := sess.GetSession(ctx, h.sessionStore, h.CustomLogger)
	if err != nil {
		return tadapter.Render(ctx, components.Notification("Ошибка получения сессии", components.NotificationError), http.StatusInternalServerError)
	}
	session.Set("email", f.Email)
	session.Set("userId", user.Id)
	session.Set("name", user.Name)
	err = sess.SaveSession(session, h.CustomLogger)
	if err != nil {
		return tadapter.Render(ctx, components.Notification("Ошибка сохранения сессии", components.NotificationError), http.StatusInternalServerError)
	}
	ctx.Response().Header.Add("Hx-Redirect", "/")
	return ctx.Redirect("/", http.StatusOK)
}

func (h ApiHandler) logout(ctx *fiber.Ctx) error {
	session, err := sess.GetSession(ctx, h.sessionStore, h.CustomLogger)
	if err != nil {
		return tadapter.Render(ctx, components.Notification("Ошибка получения сессии", components.NotificationError), http.StatusInternalServerError)
	}
	err = session.Destroy()
	if err != nil {
		return tadapter.Render(ctx, components.Notification("Ошибка очистки сессии", components.NotificationError), http.StatusInternalServerError)
	}
	ctx.Response().Header.Add("Hx-Redirect", "/")
	return ctx.Redirect("/", http.StatusOK)
}
