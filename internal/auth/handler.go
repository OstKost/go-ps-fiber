package auth

import (
	"net/http"
	"ostkost/go-ps-fiber/pkg/tadatper"
	"ostkost/go-ps-fiber/pkg/validator"
	"ostkost/go-ps-fiber/views/components"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog"
)

type AuthHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
	store        *session.Store
}

func NewAuthHandler(router fiber.Router, customLogger *zerolog.Logger, store *session.Store) {
	h := &AuthHandler{
		router:       router,
		customLogger: customLogger,
		store:        store,
	}
	authGroup := h.router.Group("/api")
	authGroup.Post("/signin", h.signIn)
	authGroup.Get("/signout", h.signOut)
}

func (h AuthHandler) signIn(ctx *fiber.Ctx) error {
	form := SignInForm{
		Email:    ctx.FormValue("email"),
		Password: ctx.FormValue("password"),
	}
	// Validation
	errors := validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: form.Email, Message: "Неправильный email"},
		&validators.StringIsPresent{Name: "Password", Field: form.Password, Message: "Не заполнен пароль"},
		&validators.StringLengthInRange{Name: "password", Field: form.Password, Min: 6, Max: 20, Message: "Пароль должен быть от 6 до 20 символов"},
	)
	var component templ.Component
	if len(errors.Errors) > 0 {
		msg := validator.FormatErrors(errors)
		component = components.Notification(msg, components.NotificationError)
		h.customLogger.Error().Msg("Validation: Sign in form errors")
		return tadatper.Render(ctx, component, http.StatusBadRequest)
	}
	// Db check password
	// err := h.repository.Create(form)
	// if err != nil {
	// 	h.customLogger.Error().Err(err).Msg("DB: Failed to create vacancy")
	// 	return tadatper.Render(ctx, components.Notification("Ошибка на сервере при создании вакансии", components.NotificationError), http.StatusBadRequest)
	// }
	if form.Email == "a@a.com" && form.Password == "123456" { // TODO: сделать нормальную авторизацию
		sess, err := h.store.Get(ctx)
		if err != nil {
			h.customLogger.Panic().Msg(err.Error())
		}
		sess.Set("email", form.Email)
		err = sess.Save()
		if err != nil {
			h.customLogger.Panic().Msg(err.Error())
		}
		ctx.Response().Header.Add("Hx-Redirect", "/")
		return ctx.Redirect("/", http.StatusOK)
	} else {
		msg := "Неверный email или пароль"
		component = components.Notification(msg, components.NotificationError)
		return tadatper.Render(ctx, component, http.StatusUnauthorized)
	}
}

func (h AuthHandler) signOut(ctx *fiber.Ctx) error {
	sess, err := h.store.Get(ctx)
	if err != nil {
		h.customLogger.Panic().Msg(err.Error())
	}
	sess.Delete("email")
	err = sess.Save()
	if err != nil {
		h.customLogger.Panic().Msg(err.Error())
	}
	ctx.Response().Header.Add("Hx-Redirect", "/")
	return ctx.Redirect("/", http.StatusOK)
}
