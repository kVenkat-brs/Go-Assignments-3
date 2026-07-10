package models

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	Id bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
	Password string `bson:"password" json:"-"`
	X_API_key string `bson:"api_key" json:"api_key"`
}