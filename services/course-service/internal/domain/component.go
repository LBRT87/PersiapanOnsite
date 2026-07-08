package domain

type Component struct {
	ID string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ModuleID string `gorm:"type:uuid"`
	Title string 
}