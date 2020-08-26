package models

type Book struct {
	ID				int		`json:"id,primary_key"`
	Title			string	`json:"title"`
	YearOfPublish	int		`json:"yearOfPublish"`
	AuthorID		int		`json:"authorId"`
}