package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/ebina4yaka/gqlgen-api-example/graph/model"
	"go.mongodb.org/mongo-driver/bson"
)

const databaseName = "example"
const collectionName = "post"

func CreatePost(post *model.Post) {
	ctx, cancel := GetContext()
	defer cancel()
	client := GetClient(ctx)
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Print(err)
		}
	}()
	col := client.Database(databaseName).Collection(collectionName)
	_, err := col.InsertOne(ctx, post)
	if err != nil {
		log.Print(err)
	}
}

func UpdatePost(id string, votes *int) {
	ctx, cancel := GetContext()
	defer cancel()
	client := GetClient(ctx)
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Print(err)
		}
	}()
	col := client.Database(databaseName).Collection(collectionName)
	filter := bson.D{{Key: "id", Value: id}}
	update := bson.D{{Key: "$set",
		Value: bson.D{
			{Key: "votes", Value: votes},
		},
	}}
	_, err := col.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Print(err)
	}
}

func FindPost(id string) *model.Post {
	ctx, cancel := GetContext()
	defer cancel()
	client := GetClient(ctx)
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Print(err)
		}
	}()
	col := client.Database(databaseName).Collection(collectionName)
	filter := bson.D{{Key: "id", Value: id}}
	var result *model.Post
	err := col.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Print(err)
	}
	return result
}

func CountPosts() int64 {
	ctx, cancel := GetContext()
	defer cancel()
	client := GetClient(ctx)
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Print(err)
		}
	}()
	col := client.Database(databaseName).Collection(collectionName)
	count, err := col.CountDocuments(ctx, bson.D{})
	if err != nil {
		log.Print(err)
	}
	return count
}

func AllPosts(limit int64, skip int64, sort int64) []*model.Post {
	ctx, cancel := GetContext()
	defer cancel()
	client := GetClient(ctx)
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Print(err)
		}
	}()
	col := client.Database(databaseName).Collection(collectionName)
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "id", Value: sort}}).SetSkip(skip).SetLimit(limit)
	cur, err := col.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		log.Print(err)
	}
	var result = make([]*model.Post, 0)
	for cur.Next(context.Background()) {
		var post model.Post
		err = cur.Decode(&post)
		if err != nil {
			log.Print(err)
		}
		result = append(result, &post)
	}
	return result
}
