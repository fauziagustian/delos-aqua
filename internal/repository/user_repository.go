package repository

import (
	"fmt"

	"github.com/fauziagustian/delos-aqua/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]*models.User, error)
	FindById(id int) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Save(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(c *gorm.DB) UserRepository {
	return &userRepository{
		db: c,
	}
}

func (r *userRepository) FindAll() ([]*models.User, error) {
	var users []*models.User

	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *userRepository) FindById(id int) (*models.User, error) {
	var user *models.User

	err := r.db.Where("user_id =?", id).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user *models.User

	fmt.Println("ini email", email)

	err := r.db.Debug().Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) Save(user *models.User) (*models.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) Update(user *models.User) (*models.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
