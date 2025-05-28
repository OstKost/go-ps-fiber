package news

import (
	"context"
	"fmt"
	"log/slog"
	"ostkost/go-ps-hw-fiber/pkg/types"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type NewsRepository struct {
	dbpool *pgxpool.Pool
	logger *slog.Logger
}

func NewNewsRepository(dbpool *pgxpool.Pool, logger *slog.Logger) *NewsRepository {
	return &NewsRepository{
		dbpool: dbpool,
		logger: logger,
	}
}

func (r *NewsRepository) Create(form types.PostNewsForm, userId int) error {
	now := time.Now().UTC()
	query := `INSERT INTO news.news (
			title, 
			preview, 
			text, 
			user_id,
			categories,
			keywords, 
			created_at, 
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.dbpool.Exec(context.Background(), query, form.Title, form.Preview, form.Text, userId, "test", "test", now, now)
	if err != nil {
		r.logger.Error(fmt.Sprintf("Failed to create news: %s", err.Error()))
		return err
	}
	return nil
}
