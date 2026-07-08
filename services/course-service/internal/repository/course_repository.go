package repository

import (
	"context"

	"github.com/LBRT87/PersiapanOnsite/services/course-service/internal/domain"
	"gorm.io/gorm"
)

type PgCourseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) domain.CourseRepository {
	return &PgCourseRepository{db: db}
}

func (r *PgCourseRepository) FindByID(c context.Context, id string) (*domain.Course, error){
	var doc domain.Course
	err := r.db.WithContext(c).First(&doc,"id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &doc, nil
}


func (r *PgCourseRepository) Create(c context.Context, course *domain.Course) error{
	return r.db.WithContext(c).Create(&course).Error
}