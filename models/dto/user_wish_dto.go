package dto

type UserWishList struct {
	ID    			int 			`json:"id"`
	UserID 			int 			`json:"user_id"`
	Name 				string 		`json:"name"`
	Price    		float32 	`json:"price"`
	Location 		string 		`json:"location"`
	IsDone 			bool 			`json:"is_done"`
}

type UserWishStatus struct {
	IsDone    	bool `json:"is_done" validate:"required"`
}
