package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	_generated "github.com/justjundana/go-graphql/graph/generated"
	_model "github.com/justjundana/go-graphql/graph/model"
	_models "github.com/justjundana/go-graphql/models"
	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input *_model.NewUser) (*_models.User, error) {
	user := _models.User{}
	user.Name = input.Name
	user.Email = input.Email
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	user.Password = string(passwordHash)

	createUser, err := r.userRepository.CreateUser(user)
	return &createUser, err
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id *int, input *_model.NewUser) (*_models.User, error) {
	user, err := r.userRepository.GetUser(*id)
	if err != nil {
		return nil, errors.New("not found")
	}

	user.Name = input.Name
	user.Email = input.Email
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	user.Password = string(passwordHash)

	updateUser, err := r.userRepository.UpdateUser(user)
	return &updateUser, err
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id *int) (*_models.User, error) {
	user, err := r.userRepository.GetUser(*id)
	if err != nil {
		return nil, errors.New("not found")
	}

	deleteUser, err := r.userRepository.DeleteUser(user)
	return &deleteUser, err
}

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

// Mutation returns _generated.MutationResolver implementation.
func (r *Resolver) Mutation() _generated.MutationResolver { return &mutationResolver{r} }

// Query returns _generated.QueryResolver implementation.
func (r *Resolver) Query() _generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
