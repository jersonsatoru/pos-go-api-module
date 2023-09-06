package database

import (
	"github.com/jersonsatoru/pos-go-api-module/internal/entities"
	"gorm.io/gorm"
)

type UserPostgresRepository struct {
	DB *gorm.DB
}

func NewUserPostgresRepository(db *gorm.DB) *UserPostgresRepository {
	return &UserPostgresRepository{
		DB: db,
	}
}

func (r *UserPostgresRepository) Create(user *entities.User) error {
	return r.DB.Create(user).Error
}

func (r *UserPostgresRepository) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
