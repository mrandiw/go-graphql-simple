package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.57

import (
	"context"

	"github.com/mrandiw/go-graphql-simple/graph/model"
	"gorm.io/gorm"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, name string, email string, address string, phone string) (*model.User, error) {
	var user model.User
	err := r.DB.Create(&model.User{Name: name, Email: email, Address: address, Phone: phone}).Error
	if err != nil {
		return nil, err
	}

	// Retrieve the newly created user to return
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) GetUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// User is the resolver for the user field.
func (r *queryResolver) GetUserDetail(ctx context.Context, id int) (*model.User, error) {
	var user model.User
	err := r.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Return nil if no user is found
		}
		return nil, err // Handle other errors
	}
	return &user, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
