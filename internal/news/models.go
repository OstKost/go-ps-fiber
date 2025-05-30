package news

import "time"

type NewsArticle struct {
	Id         int       `db:"id"`
	Title      string    `db:"title"`
	Preview    string    `db:"preview"`
	Text       string    `db:"text"`
	UserId     int       `db:"user_id"`
	Categories string    `db:"categories"`
	Keywords   string    `db:"keywords"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	DeletedAt  time.Time `db:"deleted_at"`
}
