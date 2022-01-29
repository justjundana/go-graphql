// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type LoginResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type NewBook struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Publisher   string `json:"publisher"`
	Status      bool   `json:"status"`
}

type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
