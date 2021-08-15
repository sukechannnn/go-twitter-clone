package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sukechannnn/go-twitter-clone/graph/generated"
	"github.com/sukechannnn/go-twitter-clone/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	if err := r.DB.Find(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &model.User{
		ID:         user.ID,
		Email:      user.Email,
		ScreenID:   user.ScreenID,
		ScreenName: user.ScreenName,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
