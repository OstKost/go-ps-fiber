package news

import (
	"context"
	"fmt"
	"log/slog"
	"ostkost/go-ps-hw-fiber/pkg/types"
	"strconv"
	"strings"
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

func (r *NewsRepository) Find(limit, offset int, category, keyword string) ([]NewsArticle, error) {
	var query strings.Builder
	args := make([]interface{}, 0, 4)
	argPos := 1

	query.WriteString(`SELECT 
		id, title, preview, text, user_id,
		categories, keywords, created_at, updated_at 
	FROM news.news`)

	// Добавляем условия WHERE только если есть параметры
	whereAdded := false

	if category != "" {
		query.WriteString(" WHERE categories LIKE $")
		query.WriteString(strconv.Itoa(argPos))
		args = append(args, "%"+category+"%")
		argPos++
		whereAdded = true
	}

	if keyword != "" {
		if whereAdded {
			query.WriteString(" OR keywords LIKE $")
		} else {
			query.WriteString(" WHERE keywords LIKE $")
			whereAdded = true
		}
		query.WriteString(strconv.Itoa(argPos))
		args = append(args, "%"+keyword+"%")
		argPos++
	}

	query.WriteString(" ORDER BY created_at DESC LIMIT $")
	query.WriteString(strconv.Itoa(argPos))
	args = append(args, limit)
	argPos++

	query.WriteString(" OFFSET $")
	query.WriteString(strconv.Itoa(argPos))
	args = append(args, offset)

	rows, err := r.dbpool.Query(context.Background(), query.String(), args...)
	if err != nil {
		r.logger.Error(fmt.Sprintf("Failed to find news: %s", err.Error()))
		return nil, err
	}
	news := []NewsArticle{}
	for rows.Next() {
		var n NewsArticle
		err = rows.Scan(
			&n.Id,
			&n.Title,
			&n.Preview,
			&n.Text,
			&n.UserId,
			&n.Categories,
			&n.Keywords,
			&n.CreatedAt,
			&n.UpdatedAt,
		)
		if err != nil {
			r.logger.Error(fmt.Sprintf("Failed to scan news: %s", err.Error()))
			return nil, err
		}
		news = append(news, n)
	}
	return news, nil
}
