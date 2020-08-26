package models

type Author struct {
	ID				int		`json:"id,primary_key"`
	Name			string	`json:"name"`
}