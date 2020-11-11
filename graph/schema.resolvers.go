package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/ebina4yaka/gqlgen-api-example/db"
	"github.com/ebina4yaka/gqlgen-api-example/graph/generated"
	"github.com/ebina4yaka/gqlgen-api-example/graph/model"
)

func (r *mutationResolver) CreatePost(ctx context.Context, title string, url string) (*model.Post, error) {
	count := db.CountPosts()
	post := model.Post{
		ID:        fmt.Sprintf("%d", count+1),
		Title:     title,
		URL:       url,
		Votes:     0,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	db.CreatePost(&post)
	return &post, nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, id string, votes *int) (*model.Post, error) {
	db.UpdatePost(id, votes)
	return db.FindPost(id), nil
}

func (r *queryResolver) AllPosts(ctx context.Context, orderBy *model.OrderBy, first int, skip int) ([]*model.Post, error) {
	sort := func() int64 {
		if orderBy != nil && *orderBy == "createdAt_DESC" {
			return -1
		}
		return 1
	}()
	posts := db.AllPosts(int64(first), int64(skip), sort)
	return posts, nil
}

func (r *queryResolver) AllPostsMeta(ctx context.Context) (*model.PostsMeta, error) {
	postsMeta := model.PostsMeta{Count: int(db.CountPosts())}
	return &postsMeta, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
