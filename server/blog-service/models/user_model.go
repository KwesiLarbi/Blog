package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID				primitive.ObjectID 	`json:"_id,omitempty"`
	Name 			string 				`json:"name,omitempty" validate:"required"`
	Email			string 				`json:"email,omitempty" validate:"required"`
	Password		string 				`json:"password,omitempty" validate:"required"`
	CreationTime	time.Time			`json:"reation_time,omitempty" validate:"required"` 	
}