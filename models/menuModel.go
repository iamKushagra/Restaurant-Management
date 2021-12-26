package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty" validate:"required,min=3,max=100"`
	Category   string             `json:"category,omitempty" bson:"category,omitempty" validate:"required,min=3,max=100"`
	Start_Date *time.Time         `json:"start_date" bson:"start_date"`
	End_Date   *time.Time         `json:"end_date" bson:"end_date"`
	Created_at time.Time          `json:"created_at" bson:"created_at"`
	Updated_at time.Time          `json:"updated_at" bson:"updated_at"`
	Menu_id    string             `json:"menu_id" bson:"menu_id"`
}
