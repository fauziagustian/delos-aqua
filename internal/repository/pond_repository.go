package repository

import (
	"github.com/fauziagustian/delos-aqua/internal/dto"
	"github.com/fauziagustian/delos-aqua/internal/models"
	"gorm.io/gorm"
)

type PondRepository interface {
	FindAll(query *dto.RequestQuery) ([]*models.Pond, error)
	Count() (int64, error)
	Save(pond *models.Pond) (*models.Pond, error)
	FindByName(name string) (*models.Pond, error)
	FindById(pondId int) (*models.Pond, error)
	Update(pond *models.Pond) (*models.Pond, error)
	Delete(pondId int) (*models.Pond, error)
}

type pondRepository struct {
	db *gorm.DB
}

func NewPondRepository(c *gorm.DB) PondRepository {
	return &pondRepository{
		db: c,
	}
}

func (r *pondRepository) FindAll(query *dto.RequestQuery) ([]*models.Pond, error) {
	var pond []*models.Pond

	offset := (query.Page - 1) * query.Limit
	orderBy := query.SortBy + " " + query.Sort
	queryBuider := r.db.Limit(query.Limit).Offset(offset).Order(orderBy)
	err := queryBuider.Where("name ILIKE ?", "%"+query.Search+"%").Find(&pond).Error
	if err != nil {
		return pond, err
	}

	return pond, nil
}

func (r *pondRepository) Count() (int64, error) {
	var total int64
	db := r.db.Model(&models.Pond{}).Count(&total)

	if db.Error != nil {
		return 0, db.Error
	}

	return total, nil
}

func (r *pondRepository) Save(pond *models.Pond) (*models.Pond, error) {
	err := r.db.Create(&pond).Error
	if err != nil {
		return pond, err
	}
	return pond, nil
}

func (r *pondRepository) FindByName(name string) (*models.Pond, error) {
	var pond *models.Pond

	err := r.db.Where("name = ?", name).Find(&pond).Error
	if err != nil {
		return pond, err
	}

	return pond, nil
}

func (r *pondRepository) FindById(pondId int) (*models.Pond, error) {
	var pond *models.Pond

	err := r.db.Where("pond_id = ?", pondId).Find(&pond).Error
	if err != nil {
		return pond, err
	}

	return pond, nil
}

func (r *pondRepository) Update(pond *models.Pond) (*models.Pond, error) {
	err := r.db.Model(&pond).Where("pond_id = ?", pond.PondId).Updates(pond).Error
	if err != nil {
		return pond, err
	}

	return pond, nil
}

func (r *pondRepository) Delete(pondId int) (*models.Pond, error) {
	pond := &models.Pond{}

	err := r.db.Model(&pond).Where("pond_id = ?", pondId).Delete(&pond).Error
	if err != nil {
		return pond, err
	}

	return pond, nil
}
