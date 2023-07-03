package Dto

type BusinessDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
	CoverImage  string `json:"cover_image"`
	Website     string `json:"website"`
	Location    string `json:"location"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
}
