package news

import (
	"context"
	"fmt"
	"log/slog"
	"ostkost/go-ps-hw-fiber/pkg/types"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
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
	// Инициализируем построитель запросов
	queryBuilder := sq.Select(
		"id",
		"title",
		"preview",
		"text",
		"user_id",
		"categories",
		"keywords",
		"created_at",
		"updated_at",
	).From("news.news")

	// Добавляем условия поиска для массивов
	if category != "" {
		queryBuilder = queryBuilder.Where("? = ANY(categories)", category)
	}
	if keyword != "" {
		queryBuilder = queryBuilder.Where("? = ANY(keywords)", keyword)
	}

	// Добавляем сортировку и пагинацию
	queryBuilder = queryBuilder.
		OrderBy("created_at DESC").
		Limit(uint64(limit)).
		Offset(uint64(offset))

	// Генерируем SQL и аргументы
	sql, args, err := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		r.logger.Error(fmt.Sprintf("Failed to build query: %s", err.Error()))
		return nil, err
	}

	// Выполняем запрос
	rows, err := r.dbpool.Query(context.Background(), sql, args...)
	if err != nil {
		r.logger.Error(fmt.Sprintf("Failed to find news: %s", err.Error()))
		return nil, err
	}

	// Обрабатываем результаты
	news := []NewsArticle{}
	for rows.Next() {
		var n NewsArticle
		err = rows.Scan(
			&n.Id,
			&n.Title,
			&n.Preview,
			&n.Text,
			&n.UserId,
			pq.Array(&n.Categories), // Используем pq.Array для сканирования массивов
			pq.Array(&n.Keywords),   // Используем pq.Array для сканирования массивов
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
