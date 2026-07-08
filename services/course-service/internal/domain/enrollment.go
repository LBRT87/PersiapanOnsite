package domain

import "time"

type Enrollment struct {
	ID        string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID    string `gorm:"type:uuid"`
	CourseID  string `gorm:"type:uuid"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}