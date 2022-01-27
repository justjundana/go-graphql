package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	_generated "github.com/justjundana/go-graphql/graph/generated"
	_models "github.com/justjundana/go-graphql/models"
)

func (r *queryResolver) GetUsers(ctx context.Context) ([]*_models.User, error) {
	responseData, err := r.userRepository.GetUsers()
	if err != nil {
		return nil, errors.New("not found")
	}

	userResponseData := []*_models.User{}

	for _, v := range responseData {
		userResponseData = append(userResponseData, &_models.User{ID: v.ID, Name: v.Name, Email: v.Email, Password: v.Password})
	}

	return userResponseData, nil
}

func (r *queryResolver) GetUser(ctx context.Context, id *int) (*_models.User, error) {
	responseData, err := r.userRepository.GetUser(*id)
	if err != nil {
		return nil, errors.New("not found")
	}

	return &responseData, nil
}

func (r *queryResolver) GetBooks(ctx context.Context) ([]*_models.Book, error) {
	responseData, err := r.bookRepository.GetBooks()
	if err != nil {
		return nil, errors.New("not found")
	}

	bookData := []*_models.Book{}

	for _, v := range responseData {
		bookData = append(bookData, &_models.Book{ID: v.ID, Title: v.Title, Description: v.Description, Author: v.Author, Publisher: v.Publisher, Status: v.Status})
	}

	return bookData, nil
}

// Query returns _generated.QueryResolver implementation.
func (r *Resolver) Query() _generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
