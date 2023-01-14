package message

import (
	"context"
	"time"

	repo "MessagingServer/repo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageRepo struct {
	db  *mongo.Database
	ctx context.Context
}

func NewMessageRepo(ctx context.Context, db *mongo.Database) repo.IMessageRepo {
	return &MessageRepo{ctx: ctx, db: db}
}

func (r *MessageRepo) AddMessage(userId string, data string) (err error) {

	currentDate := time.Now()
	collection := r.db.Collection("message")
	_, err = collection.InsertOne(r.ctx, bson.D{{"user_id", userId}, {"data", data}, {"created", currentDate}})

	if err != nil {
		return err
	}

	return nil
}
