package users

import (
	"net/http"
	"ostkost/go-ps-hw-fiber/pkg/types"
	"ostkost/go-ps-hw-fiber/pkg/validator"

	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	router fiber.Router
	repo   *UserRepository
}

func NewUserHandler(router fiber.Router, repo *UserRepository) {
	h := &UserHandler{
		router: router,
		repo:   repo,
	}
	usersGroup := h.router.Group("/users")
	usersGroup.Post("/", h.createUser)
	usersGroup.Get("/:email", h.getByEmail)
}

func (h UserHandler) createUser(ctx *fiber.Ctx) error {
	form := types.RegisterForm{
		Name:     ctx.FormValue("name"),
		Email:    ctx.FormValue("email"),
		Password: ctx.FormValue("password"),
	}
	err := h.repo.Create(form)
	if err != nil {
		ctx.Status(http.StatusCreated)
		return ctx.SendString(err.Error())
	}
	ctx.Status(http.StatusCreated)
	return ctx.SendString(form.Email)
}

func (h UserHandler) getByEmail(ctx *fiber.Ctx) error {
	email := ctx.Params("email")
	errors := validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: email, Message: "Неправильный email"},
	)
	if len(errors.Errors) > 0 {
		msg := validator.FormatErrors(errors)
		ctx.Status(http.StatusBadRequest)
		return ctx.SendString(msg)
	}
	user := h.repo.GetByEmail(email)
	if user == nil {
		ctx.Status(http.StatusNotFound)
		return ctx.SendString("Пользователь не найден")
	}
	return ctx.JSON(user)
}
