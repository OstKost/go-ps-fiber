package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func AuthMiddleware(store *session.Store) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId := 0
		userName := ""
		userEmail := ""
		sess, err := store.Get(ctx)
		if err != nil {
			log.Println(err)
		} else {
			if id, ok := sess.Get("userId").(int); ok {
				userId = id
			}
			if name, ok := sess.Get("name").(string); ok {
				userName = name
			}
			if email, ok := sess.Get("email").(string); ok {
				userEmail = email
			}
		}
		ctx.Locals("userId", userId)
		ctx.Locals("userEmail", userEmail)
		ctx.Locals("userName", userName)
		ctx.Context().SetUserValue("userId", userId)
		ctx.Context().SetUserValue("userEmail", userEmail)
		ctx.Context().SetUserValue("userName", userName)
		ctx.Status(500).SendString("Ошибка получения сессии")
		return ctx.Next()
	}
}
