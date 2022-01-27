package user

import (
	_models "github.com/justjundana/go-graphql/models"
)

type UserInterface interface {
	GetUsers() ([]_models.User, error)
	getUser(id int) (_models.User, error)
}
