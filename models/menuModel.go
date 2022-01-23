package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	Name       string             `json:"name" bson:"name" validate:"required,min=3,max=100"`
	Category   string             `json:"category" bson:"category" validate:"required,min=3,max=100"`
	Start_Date *time.Time         `json:"start_date" bson:"start_date"`
	End_Date   *time.Time         `json:"end_date" bson:"end_date"`
	Created_at time.Time          `json:"created_at" bson:"created_at"`
	Updated_at time.Time          `json:"updated_at" bson:"updated_at"`
	Menu_id    string             `json:"menu_id" bson:"menu_id"`
}