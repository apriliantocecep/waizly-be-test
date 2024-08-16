package repository

import (
	"gorm.io/gorm"
	"waizly/internal/entity"
)

type UserRepository struct {
	Repository[entity.User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindByEmail(db *gorm.DB, user *entity.User, email string) error {
	return db.Where("email = ?", email).First(user).Error
}

func (r *UserRepository) FindByCredential(db *gorm.DB, user *entity.User, email, password string) error {
	return db.Where("email = ? AND password = ?", email, password).First(user).Error
}
