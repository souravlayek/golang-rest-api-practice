package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Isbn   string             `json:"isbn" bson:"isbn"`
	Title  string             `json:"title" bson:"title"`
	Author *Author            `json:"author" bson:"author,omitempty"`
}
