package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/ebina4yaka/gqlgen-api-example/graph/generated"
	"github.com/ebina4yaka/gqlgen-api-example/graph/model"
)

func (r *mutationResolver) CreatePost(ctx context.Context, title string, url string) (*model.Post, error) {
	post := model.Post{
		ID:        fmt.Sprintf("%d", len(posts)+1),
		Title:     title,
		URL:       url,
		Votes:     0,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	posts = append(posts, &post)
	return &post, nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, id string, votes *int) (*model.Post, error) {
	if votes == nil {
		return nil, nil
	}
	i, _ := strconv.Atoi(id)
	posts[i-1].Votes = *votes
	return posts[i-1], nil
}

func (r *queryResolver) AllPosts(ctx context.Context, orderBy *model.OrderBy, first int, skip int) ([]*model.Post, error) {
	if skip > len(posts) {
		skip = len(posts)
	}
	if (skip + first) > len(posts) {
		first = len(posts) - skip
	}
	sortedPosts := make([]*model.Post, len(posts))
	copy(sortedPosts, posts)
	if orderBy != nil && *orderBy == "createdAt_DESC" {
		sort.SliceStable(sortedPosts, func(i, j int) bool {
			return sortedPosts[i].CreatedAt > sortedPosts[j].CreatedAt
		})
	}
	slicePosts := sortedPosts[skip : skip+first]
	return slicePosts, nil
}

func (r *queryResolver) AllPostsMeta(ctx context.Context) (*model.PostsMeta, error) {
	postsMeta := model.PostsMeta{Count: len(posts)}
	return &postsMeta, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var posts = make([]*model.Post, 0)
