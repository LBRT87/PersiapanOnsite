package domain

type Module struct {
	ID string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CourseID string `gorm:"type:uuid"`
	Title string
}