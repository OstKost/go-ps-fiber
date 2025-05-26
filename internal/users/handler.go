package users

import (
	"net/http"
	"ostkost/go-ps-hw-fiber/pkg/types"

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
	usersGroup.Get("/", h.getByEmail)
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
	user := h.repo.GetByEmail(email)
	if user == nil {
		return ctx.SendStatus(http.StatusNotFound)
	}
	return ctx.JSON(user)
}
