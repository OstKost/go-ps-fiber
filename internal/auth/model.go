package auth

type SignInForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
