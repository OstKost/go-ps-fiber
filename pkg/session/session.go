package session

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	s "github.com/gofiber/fiber/v2/middleware/session"
)

func GetSession(ctx *fiber.Ctx, sessionStore *s.Store, logger *slog.Logger) (*s.Session, error) {
	session, err := sessionStore.Get(ctx)
	if err != nil {
		logger.Error("Failed to get session")
		return nil, err
	}
	return session, nil
}

func SaveSession(session *s.Session, logger *slog.Logger) error {
	err := session.Save()
	if err != nil {
		logger.Error("Failed to save session")
		return err
	}
	return nil
}
