package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id"`
	Invoice_id       string             `json:"invoice_id" bson:"invoice_id"`
	Order_id         string             `json:"order_id" bson:"order_id"`
	Payment_method   *string            `json:"payment_method" bson:"payment_method" validate:"eq=CARD|eq=CASH|eq="`
	Payment_status   *string            `json:"payment_status" bson:"payment_status" validate:"eq=PAID|eq=UNPAID"`
	Payment_due_date time.Time          `json:"payment_due_date" bson:"payment_due_date"`
	Created_at       time.Time          `json:"created_at" bson:"created_at"`
	Updated_at       time.Time          `json:"updated_at" bson:"updated_at"`
}
