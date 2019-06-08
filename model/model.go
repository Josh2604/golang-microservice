package model

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	UserId	   int 						`bson:"user_id" json:"user_id"`
	Name       string        `bson:"name" json:"name"`
	Apells     string       `bson:"apells" json:"apells"`
	Age        int           `bson:"age" json:"age"`
	Email      string        `bson:"email" json:"email"`
}