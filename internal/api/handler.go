package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"ostkost/go-ps-hw-fiber/internal/news"
	"ostkost/go-ps-hw-fiber/internal/users"
	sess "ostkost/go-ps-hw-fiber/pkg/session"
	"ostkost/go-ps-hw-fiber/pkg/tadapter"
	"ostkost/go-ps-hw-fiber/pkg/types"
	"ostkost/go-ps-hw-fiber/pkg/validator"
	"ostkost/go-ps-hw-fiber/views/components"
	"ostkost/go-ps-hw-fiber/views/widgets"
	"strconv"

	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	s "github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
)

type ApiHandler struct {
	router       fiber.Router
	logger       *slog.Logger
	userRepo     *users.UserRepository
	newsRepo     *news.NewsRepository
	sessionStore *s.Store
}

type ApiHandlerProps struct {
	Router       fiber.Router
	Logger       *slog.Logger
	UserRepo     *users.UserRepository
	NewsRepo     *news.NewsRepository
	SessionStore *s.Store
}

func NewApiHandler(props ApiHandlerProps) {
	h := &ApiHandler{
		router:       props.Router,
		logger:       props.Logger,
		userRepo:     props.UserRepo,
		newsRepo:     props.NewsRepo,
		sessionStore: props.SessionStore,
	}
	apiGroup := h.router.Group("/api")
	apiGroup.Post("/register", h.register)
	apiGroup.Post("/login", h.login)
	apiGroup.Post("/logout", h.logout)
	apiGroup.Post("/news", h.createNews)
	apiGroup.Get("/news", h.findNews)
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
	session, err := sess.GetSession(ctx, h.sessionStore, h.logger)
	if err != nil {
		return tadapter.Render(ctx, components.Notification("Ошибка получения сессии", components.NotificationError), http.StatusInternalServerError)
	}
	session.Set("email", f.Email)
	session.Set("userId", user.Id)
	session.Set("name", user.Name)
	err = sess.SaveSession(session, h.logger)
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
	// Validation
	errors := validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: f.Email, Message: "Неправильный email"},
		&validators.StringIsPresent{Name: "Password", Field: f.Password, Message: "Не заполнен пароль"},
		&validators.StringLengthInRange{Name: "password", Field: f.Password, Min: 6, Max: 20, Message: "Пароль должен быть от 6 до 20 символов"},
	)
	if len(errors.Errors) > 0 {
		msg := validator.FormatErrors(errors)
		component := components.Notification(msg, components.NotificationError)
		return tadapter.Render(ctx, component, http.StatusBadRequest)
	}
	// Get user from db
	user := h.userRepo.GetByEmail(f.Email)
	if user == nil {
		return tadapter.Render(ctx, components.Notification("Неверный email или пароль", components.NotificationError), http.StatusUnauthorized)
	}
	// Check hashed password with form password

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(f.Password))
	if err != nil {
		return tadapter.Render(ctx, components.Notification("Неверный email или пароль", components.NotificationError), http.StatusUnauthorized)
	}

	// Session
	session, err := sess.GetSession(ctx, h.sessionStore, h.logger)
	if err != nil {
		return tadapter.Render(ctx, components.Notification("Ошибка получения сессии", components.NotificationError), http.StatusInternalServerError)
	}
	session.Set("email", f.Email)
	session.Set("userId", user.Id)
	session.Set("name", user.Name)
	err = sess.SaveSession(session, h.logger)
	if err != nil {
		return tadapter.Render(ctx, components.Notification("Ошибка сохранения сессии", components.NotificationError), http.StatusInternalServerError)
	}
	ctx.Response().Header.Add("Hx-Redirect", "/")
	return ctx.Redirect("/", http.StatusOK)
}

func (h ApiHandler) logout(ctx *fiber.Ctx) error {
	session, err := sess.GetSession(ctx, h.sessionStore, h.logger)
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

func (h ApiHandler) createNews(ctx *fiber.Ctx) error {
	userId, _ := ctx.Context().Value("userId").(int)
	if userId == 0 {
		return tadapter.Render(ctx, components.Notification("Необходимо авторизоваться", components.NotificationError), http.StatusUnauthorized)
	}
	f := types.PostNewsForm{
		Title:   ctx.FormValue("title"),
		Preview: ctx.FormValue("preview"),
		Text:    ctx.FormValue("text"),
	}
	errors := validate.Validate(
		&validators.StringIsPresent{Name: "Title", Field: f.Title, Message: "Не задан заголовок"},
		&validators.StringIsPresent{Name: "Preview", Field: f.Preview, Message: "Не задано превью"},
		&validators.StringIsPresent{Name: "Text", Field: f.Text, Message: "Не задан текст"},
	)
	if len(errors.Errors) > 0 {
		msg := validator.FormatErrors(errors)
		component := components.Notification(msg, components.NotificationError)
		return tadapter.Render(ctx, component, http.StatusBadRequest)
	}
	// Db insert
	err := h.newsRepo.Create(f, userId)
	if err != nil {
		h.logger.Error(fmt.Sprintf("DB: Failed to create news: %s", err.Error()))
		return tadapter.Render(ctx, components.Notification("Ошибка на сервере при создании новости", components.NotificationError), http.StatusInternalServerError)
	}
	return tadapter.Render(ctx, components.Notification("Новость создана", components.NotificationSuccess), http.StatusCreated)
}

func (h ApiHandler) findNews(ctx *fiber.Ctx) error {
	limit := ctx.QueryInt("limit")
	offset := ctx.QueryInt("offset")
	category := ctx.Query("category")
	keyword := ctx.Query("keyword")
	if limit == 0 {
		limit = 10
	}
	if offset == 0 {
		offset = 0
	}
	news, err := h.newsRepo.Find(limit, offset, category, keyword)
	if err != nil {
		h.logger.Error(fmt.Sprintf("DB: Failed to find news: %s", err.Error()))
		return tadapter.Render(ctx, components.Notification("Ошибка на сервере при поиске новостей", components.NotificationError), http.StatusInternalServerError)
	}
	var posts []components.PostCardProps
	for _, n := range news {
		posts = append(posts, components.PostCardProps{
			Title:       n.Title,
			Description: n.Text,
			Img:         "nature.jpg",
			Username:    strconv.Itoa(n.UserId),
			AvatarImg:   "Mike.jpg",
			Date:        n.CreatedAt.Format("02.01.2006"),
		})
	}
	component := widgets.Posts(posts)
	return tadapter.Render(ctx, component, http.StatusCreated)
}
