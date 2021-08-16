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
	id, err := model.CreateUser(r.DB, input)
	if err != nil {
		return nil, err
	}
	return model.FindById(r.DB, id)
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	user := ForContext(ctx)
	if user == nil {
		return &model.User{}, fmt.Errorf("access denied")
	}
	return user, nil
}

func (r *queryResolver) FollowUsers(ctx context.Context) ([]*model.User, error) {
	user := ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	var followUsers []*model.FollowUser
	if err := r.DB.Where("user_id = ?", user.ID).Find(&followUsers).Error; err != nil {
		return nil, err
	}

	var ids []string
	for _, v := range followUsers {
		ids = append(ids, v.FollowerID)
	}
	var users []*model.User
	if err := r.DB.Where("id in ?", ids).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *queryResolver) Followers(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
