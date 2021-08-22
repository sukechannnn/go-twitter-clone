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
	userRepo := model.UserRepository{DB: r.DB}
	id, err := userRepo.Create(input)
	if err != nil {
		return nil, err
	}
	return userRepo.FindById(id)
}

func (r *mutationResolver) FollowUser(ctx context.Context, input model.NewFollowUser) (*model.FollowUser, error) {
	user := ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	fUserRepo := model.FollowUserRepository{DB: r.DB}
	id, err := fUserRepo.Create(input, user.ID)
	if err != nil {
		return nil, err
	}
	return fUserRepo.FindById(id)
}

func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	user := ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	postRepo := model.PostRepository{DB: r.DB}
	id, err := postRepo.Create(input, user.ID)
	if err != nil {
		return nil, err
	}
	return postRepo.FindById(id)
}

func (r *queryResolver) AllUsers(ctx context.Context) ([]*model.UserInfo, error) {
	user := ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	userRepo := model.UserRepository{DB: r.DB}
	return userRepo.All(user.ID)
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	user := ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	return user, nil
}

func (r *queryResolver) FollowUsers(ctx context.Context) ([]*model.User, error) {
	user := ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	fUserRepo := model.FollowUserRepository{DB: r.DB}
	var followUsers []*model.FollowUser
	followUsers, err := fUserRepo.FollowUsers(user.ID)
	if err != nil {
		return nil, err
	}

	var ids []string
	for _, v := range followUsers {
		ids = append(ids, v.FollowID)
	}
	userRepo := model.UserRepository{DB: r.DB}
	return userRepo.FindByIds(ids)
}

func (r *queryResolver) Followers(ctx context.Context) ([]*model.User, error) {
	user := ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	fUserRepo := model.FollowUserRepository{DB: r.DB}
	var followrs []*model.FollowUser
	followrs, err := fUserRepo.Followers(user.ID)
	if err != nil {
		return nil, err
	}

	var ids []string
	for _, v := range followrs {
		ids = append(ids, v.UserID)
	}
	userRepo := model.UserRepository{DB: r.DB}
	return userRepo.FindByIds(ids)
}

func (r *queryResolver) Timeline(ctx context.Context) ([]*model.PostInfo, error) {
	user := ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	fUserRepo := model.FollowUserRepository{DB: r.DB}
	var followUsers []*model.FollowUser
	followUsers, err := fUserRepo.FollowUsers(user.ID)
	if err != nil {
		return nil, err
	}

	ids := []string{user.ID}
	for _, v := range followUsers {
		ids = append(ids, v.FollowID)
	}
	postRepo := model.PostRepository{DB: r.DB}
	return postRepo.Timeline(ids)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
