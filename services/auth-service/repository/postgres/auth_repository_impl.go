package postgres

import (
	"context"
	"errors"

	"github.com/LBRT87/PersiapanOnsite/services/auth-service/entity"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) entity.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) GetByEmail(ctx context.Context, user *entity.User) entity.UserRepository {
	var user entity.User

	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error

	if erros.is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, err
}

func (r *userRepository) GetByID(ctx context.Context, id uint) (*entity.UserRepository, error) {
	var user entity.User

	err := r.db.WithContext(ctx).First(&user, id).Error

	if erros.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByGoogle(ctx context.Context, googleID string) (*entity.User, error) {
	var user entity.User

	err := r.db.WithContext(ctx).Where("google_id = ?", googleID).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
