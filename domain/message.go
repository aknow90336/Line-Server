package domain

import "time"

type Message struct {
	Message string    `json:"message" bson:"message,omitempty"`
	Created time.Time `json:"created" bson:"created,omitempty"`
}
