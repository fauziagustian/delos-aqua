package repository

import (
	"github.com/fauziagustian/delos-aqua/internal/dto"
	"github.com/fauziagustian/delos-aqua/internal/models"
	"gorm.io/gorm"
)

type FarmRepository interface {
	FindAll(query *dto.RequestQuery, userid int) ([]*models.Farm, error)
	Count() (int64, error)
	Save(farm *models.Farm, userid int) (*models.Farm, error)
	FindByName(name string) (*models.Farm, error)
	FindById(farmId int) (*models.Farm, error)
	Update(farm *models.Farm, userid int) (*models.Farm, error)
	Delete(farmId, userid int) (*models.Farm, error)
}

type farmRepository struct {
	db *gorm.DB
}

func NewFarmRepository(c *gorm.DB) FarmRepository {
	return &farmRepository{
		db: c,
	}
}

func (r *farmRepository) FindAll(query *dto.RequestQuery, userid int) ([]*models.Farm, error) {
	var farm []*models.Farm

	offset := (query.Page - 1) * query.Limit
	orderBy := query.SortBy + " " + query.Sort
	queryBuider := r.db.Limit(query.Limit).Offset(offset).Order(orderBy)
	err := queryBuider.Where("name ILIKE ?", "%"+query.Search+"%").Find(&farm).Error
	if err != nil {
		return farm, err
	}

	userAgent := &models.UserAgents{}
	userAgent.MethodUrl = "GET /farm"
	userAgent.UserId = userid
	SaveUserAgent(r.db, userAgent)

	return farm, nil
}

func (r *farmRepository) Count() (int64, error) {
	var total int64
	db := r.db.Model(&models.Farm{}).Count(&total)

	if db.Error != nil {
		return 0, db.Error
	}

	return total, nil
}

func (r *farmRepository) Save(farm *models.Farm, userid int) (*models.Farm, error) {
	err := r.db.Create(&farm).Error
	if err != nil {
		return farm, err
	}

	userAgent := &models.UserAgents{}
	userAgent.MethodUrl = "POST /farm"
	userAgent.UserId = userid

	SaveUserAgent(r.db, userAgent)

	return farm, nil
}

func (r *farmRepository) FindByName(name string) (*models.Farm, error) {
	var farm *models.Farm

	err := r.db.Where("name = ?", name).Find(&farm).Error
	if err != nil {
		return farm, err
	}

	return farm, nil
}

func (r *farmRepository) FindById(farmId int) (*models.Farm, error) {
	var farm *models.Farm

	err := r.db.Where("farm_id = ?", farmId).Find(&farm).Error
	if err != nil {
		return farm, err
	}

	return farm, nil
}

func (r *farmRepository) Update(farm *models.Farm, userid int) (*models.Farm, error) {
	err := r.db.Model(&farm).Where("farm_id = ?", farm.FarmId).Update("name", farm.Name).Error
	if err != nil {
		return farm, err
	}

	userAgent := &models.UserAgents{}
	userAgent.MethodUrl = "PUT /farm"
	userAgent.UserId = userid

	SaveUserAgent(r.db, userAgent)

	return farm, nil
}

func (r *farmRepository) Delete(farmId, userid int) (*models.Farm, error) {
	farm := &models.Farm{}

	err := r.db.Model(&farm).Where("farm_id = ?", farmId).Delete(&farm).Error
	if err != nil {
		return farm, err
	}

	userAgent := &models.UserAgents{}
	userAgent.MethodUrl = "DELETE /farm"
	userAgent.UserId = userid

	SaveUserAgent(r.db, userAgent)

	return farm, nil
}
