package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func AuthMiddleware(store *session.Store) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		sess, err := store.Get(ctx)
		if err != nil {
			ctx.Status(500).SendString("Ошибка получения сессии")
			return ctx.Next()
		}
		userId := 0
		userName := ""
		userEmail := ""
		if id, ok := sess.Get("userId").(int); ok {
			userId = id
		}
		if name, ok := sess.Get("name").(string); ok {
			userName = name
		}
		if email, ok := sess.Get("email").(string); ok {
			userEmail = email
		}
		// Сохраняем данные в Locals
		ctx.Locals("userId", userId)
		ctx.Locals("userEmail", userEmail)
		ctx.Locals("userName", userName)
		// Дополнительно сохраняем в контекст для совместимости
		ctx.Context().SetUserValue("userId", userId)
		ctx.Context().SetUserValue("userEmail", userEmail)
		ctx.Context().SetUserValue("userName", userName)
		return ctx.Next()
	}
}
