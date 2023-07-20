package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Student struct {
	Id    primitive.ObjectID `json:"_id" bson:"_id"`
	Email string             `json:"email" bson:"email"`
}
