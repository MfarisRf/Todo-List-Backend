package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Checklist struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	Title  string             `bson:"title" json:"title"`
	Items  []TodoItem         `bson:"items" json:"items"`
}

type TodoItem struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name   string             `bson:"name" json:"name"`
	Status string             `bson:"status" json:"status"`
}
