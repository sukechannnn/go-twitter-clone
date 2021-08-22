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
	return model.FindUserById(r.DB, id)
}

func (r *mutationResolver) FollowUser(ctx context.Context, input model.NewFollowUser) (*model.FollowUser, error) {
	user := ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	id, err := model.CreateFollowUser(r.DB, input, user.ID)
	if err != nil {
		return nil, err
	}
	return model.FindFollowUserById(r.DB, id)
}

func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	user := ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	id, err := model.CreatePost(r.DB, input, user.ID)
	if err != nil {
		return nil, err
	}
	var post model.Post
	if err := r.DB.Find(&post, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *queryResolver) AllUsers(ctx context.Context) ([]*model.User, error) {
	user := ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	return model.AllUsers(r.DB, user.ID)
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
	var followUsers []*model.FollowUser
	if err := r.DB.Where("user_id = ?", user.ID).Find(&followUsers).Error; err != nil {
		return nil, err
	}

	var ids []string
	for _, v := range followUsers {
		ids = append(ids, v.FollowID)
	}
	var users []*model.User
	if err := r.DB.Where("id in ?", ids).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *queryResolver) Followers(ctx context.Context) ([]*model.User, error) {
	user := ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	var followrs []*model.FollowUser
	if err := r.DB.Where("follow_id = ?", user.ID).Find(&followrs).Error; err != nil {
		return nil, err
	}

	var ids []string
	for _, v := range followrs {
		ids = append(ids, v.UserID)
	}
	var users []*model.User
	if err := r.DB.Where("id in ?", ids).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *queryResolver) Timeline(ctx context.Context) ([]*model.Post, error) {
	user := ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	var followUsers []*model.FollowUser
	if err := r.DB.Where("user_id = ?", user.ID).Find(&followUsers).Error; err != nil {
		return nil, err
	}

	ids := []string{user.ID}
	for _, v := range followUsers {
		ids = append(ids, v.FollowID)
	}
	var posts []*model.Post
	if err := r.DB.Where("user_id in ?", ids).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
