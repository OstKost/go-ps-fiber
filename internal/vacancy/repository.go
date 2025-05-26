package vacancy

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type VacancyRepository struct {
	Dbpool       *pgxpool.Pool
	CustomLogger *zerolog.Logger
}

func NewVacancyRepository(dbpool *pgxpool.Pool, customLogger *zerolog.Logger) *VacancyRepository {
	return &VacancyRepository{
		Dbpool:       dbpool,
		CustomLogger: customLogger,
	}
}

func (r *VacancyRepository) Create(form PostVacancyForm) error {
	// Query
	query := `INSERT INTO vacancies (vacancy, company, industry, salary, location, email, created_at) 
			VALUES (@vacancy, @company, @industry, @salary, @location, @email, @created_at)`
	// Args
	args := pgx.NamedArgs{
		"vacancy":    form.Vacancy,
		"company":    form.Company,
		"industry":   form.Industry,
		"salary":     form.Salary,
		"location":   form.Location,
		"email":      form.Email,
		"created_at": time.Now(),
	}
	// Execute
	_, err := r.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("failed to create vacancy: %w", err)
	}
	return nil
}

func (r *VacancyRepository) GetAll(limit, offset int) ([]Vacancy, error) {
	// var vacancies []*Vacancy
	query := `SELECT id, vacancy, company, industry, salary, location, email, created_at
			FROM vacancies
			ORDER BY created_at DESC
			LIMIT @limit OFFSET @offset`
	args := pgx.NamedArgs{
		"limit":  limit,
		"offset": offset,
	}
	rows, err := r.Dbpool.Query(context.Background(), query, args)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	vacancies, err := pgx.CollectRows(rows, pgx.RowToStructByName[Vacancy])
	if err != nil {
		return nil, err
	}
	return vacancies, nil

	// Manual example
	// for rows.Next() {
	// 	var vacancy Vacancy
	// 	err = rows.Scan(&vacancy.Id, &vacancy.Vacancy, &vacancy.Company, &vacancy.Industry, &vacancy.Salary, &vacancy.Location, &vacancy.Email)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	vacancies = append(vacancies, &vacancy)
	// }
}

func (r *VacancyRepository) CountAll() int {
	query := "SELECT count(*) FROM vacancies"
	var count int
	r.Dbpool.QueryRow(context.Background(), query).Scan(&count)
	return count
}
