package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"fmt"

	graph "github.com/rsp84/ride-with-me-backend/graph/generated"
	"github.com/rsp84/ride-with-me-backend/graph/models"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, name string, email string) (*models.User, error) {
	user := &models.User{
		Name:  name,
		Email: email,
	}
	result := r.DB.Create(user)
	if result.Error != nil {
		if result.Error.Error() == `ERROR: duplicate key value violates unique constraint "idx_users_email" (SQLSTATE 23505)` {
			return nil, fmt.Errorf("email already exists")
		}
		return nil, fmt.Errorf("failed to create user: %w", result.Error)
	}
	return user, nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
