package vacancy

type PostVacancyForm struct {
	Vacancy     string `json:"vacancy"`
	Company     string `json:"company"`
	Industry    string `json:"industry"`
	Salary      string `json:"salary"`
	Location    string `json:"location"`
	Email       string `json:"email"`
	SubmitError string `json:"submitError"`
}
