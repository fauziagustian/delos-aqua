package repository

import (
	"log"

	"github.com/fauziagustian/delos-aqua/internal/dto"
	"github.com/fauziagustian/delos-aqua/internal/models"
	"gorm.io/gorm"
)

type UserAgentRepository interface {
	FindAll() ([]*dto.UserAgentResponseCount, error)
}

type userAgentRepository struct {
	db *gorm.DB
}

func NewUserAgentRepository(c *gorm.DB) UserAgentRepository {
	return &userAgentRepository{
		db: c,
	}
}

func (r *userAgentRepository) FindAll() ([]*dto.UserAgentResponseCount, error) {
	var ua []*dto.UserAgentResponseCount

	err := r.db.Raw("select method_url, count(method_url) as count, count(distinct user_id) as unique_user_agent from user_agents group by method_url").Scan(&ua).Error
	if err != nil {
		return ua, err
	}

	log.Println("Cek variabel UA", ua)

	return ua, nil
}

func SaveUserAgent(db *gorm.DB, userAgent *models.UserAgents) {
	err := db.Create(&userAgent).Error
	if err != nil {
		log.Println(err)
	}
}
