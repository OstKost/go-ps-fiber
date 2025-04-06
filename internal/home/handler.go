package home

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

type HomeHandler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	h := &HomeHandler{
		router: router,
	}
	h.router.Get("/", h.home)
}

type User struct {
	Name string
	Age  int
}

func (h HomeHandler) home(ctx *fiber.Ctx) error {
	//ctx.Hostname()
	//tmpl := template.Must(template.ParseFiles("./html/page.html"))
	//data := struct {
	//	Count int
	//	Name  string
	//}{100, "foo"}
	//var tpl bytes.Buffer
	//if err := tmpl.Execute(&tpl, data); err != nil {
	//	return fiber.NewError(fiber.StatusInternalServerError, "Template compile error")
	//}
	//ctx.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	//return ctx.Send(tpl.Bytes())

	words := []string{"foo", "bar", "baz"}
	users := []User{
		{Name: "Ivan", Age: 10},
		{Name: "Vasya", Age: 20},
		{Name: "Kostya", Age: 30},
	}

	data := fiber.Map{
		"Count":   100,
		"Name":    "foo",
		"IsAdmin": false,
		"Users":   users,
		"Words":   words,
		"ToUpper": func(s string) string {
			return strings.ToUpper(s)
		},
	}

	return ctx.Render("page", data)
}
