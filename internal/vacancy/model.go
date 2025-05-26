package vacancy

import "time"

type PostVacancyForm struct {
	Vacancy   string    `json:"vacancy"`
	Company   string    `json:"company"`
	Industry  string    `json:"industry"`
	Salary    string    `json:"salary"`
	Location  string    `json:"location"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type Vacancy struct {
	Id        int       `db:"id"`
	Vacancy   string    `db:"vacancy"`
	Company   string    `db:"company"`
	Industry  string    `db:"industry"`
	Salary    string    `db:"salary"`
	Location  string    `db:"location"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
}
