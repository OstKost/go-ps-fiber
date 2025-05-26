package users

import (
	"context"
	"fmt"
	"log/slog"
	"ostkost/go-ps-hw-fiber/pkg/types"

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
	user := r.GetByEmail(form.Email)
	if user != nil {
		return fmt.Errorf("Email already exists")
	}
	query := `INSERT INTO users (email, password, name) 
			VALUES ($1, $2, $3)`
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		r.logger.Error(fmt.Sprintf("Failed to hash password: %s", err.Error()))
		return err
	}
	_, err = r.dbpool.Exec(context.Background(), query, form.Email, string(hashedPassword), form.Name)
	if err != nil {
		r.logger.Error(fmt.Sprintf("Failed to create user: %s", err.Error()))
		return err
	}
	return nil
}

func (r *UserRepository) GetByEmail(email string) *types.User {
	query := `SELECT id, email, password, name, created_at
			FROM users
			WHERE email = $1`
	var user types.User
	row := r.dbpool.QueryRow(context.Background(), query, email)
	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.CreatedAt)
	if err != nil {
		r.logger.Error(fmt.Sprintf("Failed to get user by email: %s", err.Error()))
		return nil
	}
	return &user
}
