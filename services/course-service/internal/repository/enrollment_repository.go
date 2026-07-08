package repository

import (
	"context"
	"github.com/LBRT87/PersiapanOnsite/services/course-service/internal/domain"
	"gorm.io/gorm"
)

type PgEnrollmentRepository struct {
	db *gorm.DB
}

func (p *PgEnrollmentRepository) Create(c context.Context, enrollment *domain.Enrollment) error {
	return p.db.WithContext(c).Create(&enrollment).Error
}

func (p *PgEnrollmentRepository) FindByID(c context.Context, id string) (*domain.Enrollment, error) {
	var enr domain.Enrollment
	err := p.db.WithContext(c).First(&enr,"id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &enr, nil
}

func (p *PgEnrollmentRepository) FindByUserCourse(c context.Context, userId string, courseId string) (bool, error) {
	var enr domain.Enrollment
	err := p.db.WithContext(c).Where("user_id = ? AND lecturer_id = ?", userId, courseId).First(&enr).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewEnrollmentRepository(db *gorm.DB) domain.EnrollmentRepository {
	return &PgEnrollmentRepository{db: db}
}
