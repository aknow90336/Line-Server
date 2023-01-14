package message

import (
	"context"
	"log"
	"time"

	domain "MessagingServer/domain"
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
	_, err = collection.InsertOne(r.ctx, bson.D{{"user_id", userId}, {"message", data}, {"created", currentDate}})

	if err != nil {
		return err
	}

	return nil
}

func (r *MessageRepo) GetMessageByUserId(userId string) (record []*domain.Message, err error) {
	record = make([]*domain.Message, 0)

	collection := r.db.Collection("message")
	cur, err := collection.Find(r.ctx, bson.D{{"user_id", userId}})
	if err != nil {
		log.Println(err)
	}
	defer cur.Close(r.ctx)
	for cur.Next(r.ctx) {

		result := &domain.Message{}
		err := cur.Decode(&result)
		record = append(record, result)
		if err != nil {
			log.Println(err)
		}
	}

	if err := cur.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return record, nil
}
