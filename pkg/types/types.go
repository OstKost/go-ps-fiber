package types

type RegisterForm struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Id        int    `db:"id"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	Name      string `db:"name"`
	CreatedAt string `db:"created_at"`
}
