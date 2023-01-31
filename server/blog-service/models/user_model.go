package models

import "time"

type User struct {
	UserID 					string 		`json:"user_id"`
	Name 					string 		`json:"name"`
	Email					string 		`json:"email"`
	Password				string 		`json:"password"`
	AccountCreationDateTime	time.Time	`json:"account_creation_time"` 	
}