package users

import (
	"context"
	"fmt"
	"log/slog"
	"ostkost/go-ps-hw-fiber/pkg/types"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	dbpool *pgxpool.Pool
	logger *slog.Logger
}

func NewUserRepository(dbpool *pgxpool.Pool, customLogger *slog.Logger) *UserRepository {
	return &UserRepository{
		dbpool: dbpool,
		logger: customLogger,
	}
}

func (r *UserRepository) Create(form types.RegisterForm) error {
	query := `INSERT INTO users (email, password, name) 
			VALUES (@email, @password, @name)`
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		r.logger.Error(fmt.Sprintf("Failed to hash password: %s", err.Error()))
	}
	args := pgx.NamedArgs{
		"email":    form.Email,
		"password": string(hashedPassword),
		"name":     form.Name,
	}
	_, err = r.dbpool.Exec(context.Background(), query, args)
	if err != nil {
		r.logger.Error(fmt.Sprintf("Failed to create user: %s", err.Error()))
		return err
	}
	return nil
}

func (r *UserRepository) GetByEmail(email string) *types.User {
	query := `SELECT id, email, password, name, created_at
			FROM users
			WHERE email = @email`
	args := pgx.NamedArgs{
		"email": email,
	}
	var user types.User
	r.dbpool.QueryRow(context.Background(), query, args).
		Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.CreatedAt)
	return &user
}
