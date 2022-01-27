package models

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Publisher   string `json:"publisher"`
	Status      bool   `json:"status"`
}
