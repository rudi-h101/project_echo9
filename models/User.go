package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   				int 			`json:"id"`
	Name 				string 		`json:"name" validate:"required"`
	Email    		string 		`json:"email" gorm:"unique" validate:"required,email"`
	Password 		string 		`json:"password" validate:"required"`
	UserWish    []UserWish 

	CreatedAt   time.Time	`json:"created_at"`
  UpdatedAt   time.Time	`json:"updated_at"`
}