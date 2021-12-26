package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	First_name    *string            `json:"first_name" bson:"first_name" validate:"required,min=3,max=100"`
	Last_name     *string            `json:"last_name" bson:"last_name" validate:"required,min=1,max=100"`
	Password      *string            `json:"password" bson:"password" validate:"required,min=6"`
	Email         *string            `json:"email" bson:"email" validate:"required"`
	Avatar        *string            `json:"avatar" bson:"avatar"`
	Phone         *string            `json:"phone" bson:"phone" validate:"required,min=10"`
	Token         *string            `json:"token" bson:"token"`
	Refresh_Token *string            `json:"refresh_token" bson:"refresh_token"`
	Created_at    time.Time          `json:"created_at" bson:"created_at"`
	Updated_at    time.Time          `json:"updated_at" bson:"updated_at"`
	User_id       string             `json:"user_id" bson:"user_id"`
}
