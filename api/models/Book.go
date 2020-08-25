package models

type Book struct {
	ID int `json:"id,primary_key"`
	Title string
	YearOfPublish int
}