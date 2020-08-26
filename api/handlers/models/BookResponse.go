package apimodels

type BookResponse struct {
	ID				int					`json:"id"`
	Title			string				`json:"title"`
	YearOfPublish	int					`json:"yearOfPublish"`
	Author			BookAuthorResponse	`json:"author"`
}

type BookAuthorResponse struct {
	ID		int		`json:"id"`
	Name	string	`json:"name"`
}
