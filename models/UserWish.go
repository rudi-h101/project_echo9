package models

import (
	"time"

	"gorm.io/gorm"
)

type UserWish struct {
	gorm.Model
	ID        	int 			`json:"id"`
	UserID    	int 			`json:"user_id"`	
	Name 				string 		`json:"name" validate:"required"`
	Price    		float32 	`json:"price" validate:"required"`
	Location 		string 		`json:"location" validate:"required"`
	IsDone 			bool 			`json:"is_done" gorm:"default:false"`
	CreatedAt   time.Time	`json:"created_at"`
  UpdatedAt   time.Time	`json:"updated_at"`
}
