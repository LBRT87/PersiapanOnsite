package domain

import "time"

type Course struct {
	ID        	string `gorm:"type:uuid;primaryKey;default:gen_random_uuid"`
	LectureID 	string `json:"lecturer_id"`
	Title 	    string `json:"title"`
	Description string `json:"description"`
	Status 		string `json:"status"`
	Price     	uint `json:"price"`
	BuyerCount 	uint `json:"buyer_count"`
	CreatedAt 	time.Time `gorm:"autoCreateTime"`
}