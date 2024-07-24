package repository

import (
	"errors"

	"github.com/nuhmanudheent/hotel-booking/user-service/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(user domain.User) error
	FindByEmail(email string) (domain.User, error)
	FindByID(id uint) (domain.User, error)
	CheckUser(userID uint32) bool
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) RegisterUser(user domain.User) error {
	if err := u.db.Create(&user).Error; err != nil {
		return errors.New("email already exists")
	}
	return nil
}

func (u *userRepository) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return domain.User{}, errors.New("invalid email or password")
	}
	return user, nil
}

func (u *userRepository) FindByID(id uint) (domain.User, error) {
	var user domain.User
	if err := u.db.First(&user, id).Error; err != nil {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}
func (u *userRepository) CheckUser(userID uint32) bool {
	var user domain.User
	if err := u.db.First(&user, userID).Error; err != nil {
		return false
	}
	return true
}
