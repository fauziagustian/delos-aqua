package repository

import (
	"github.com/fauziagustian/delos-aqua/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]*models.User, error)
	FindById(id int) (*models.User, error)
	FindByName(name string) ([]*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Save(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

type URConfig struct {
	DB *gorm.DB
}

func NewUserRepository(c *URConfig) UserRepository {
	return &userRepository{
		db: c.DB,
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

	err := r.db.Where("id =?", id).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindByName(name string) ([]*models.User, error) {
	var users []*models.User

	err := r.db.Where("name ILIKE ?", "%"+name+"%").Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user *models.User

	err := r.db.Where("email = ?", email).Find(&user).Error
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
