package graph

import (
	_bookRepository "github.com/justjundana/go-graphql/repository/book"
	_userRepository "github.com/justjundana/go-graphql/repository/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userRepository _userRepository.UserInterface
	bookRepository _bookRepository.BookInterface
}

func NewResolver(ur _userRepository.UserInterface, br _bookRepository.BookInterface) *Resolver {
	return &Resolver{
		userRepository: ur,
		bookRepository: br,
	}
}
