package user

import (
	_models "github.com/justjundana/go-graphql/models"
)

type UserInterface interface {
	GetUsers() ([]_models.User, error)
	GetUser(id int) (_models.User, error)
	CreateUser(user _models.User) (_models.User, error)
	UpdateUser(user _models.User) (_models.User, error)
}
