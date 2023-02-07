package service

import (
	"log"

	"github.com/fauziagustian/delos-aqua/internal/dto"
	"github.com/fauziagustian/delos-aqua/internal/models"
	r "github.com/fauziagustian/delos-aqua/internal/repository"
	custom_error "github.com/fauziagustian/delos-aqua/pkg/customError"
)

type FarmService interface {
	GetFarm(query *dto.FarmRequestQuery) ([]*models.Farm, error)
	CountFarm() (int64, error)
	CreateFarm(query *dto.FarmRequestBody) (*models.Farm, error)
	UpdateFarm(query *dto.FarmRequestBody, id int) (*models.Farm, error)
	DeleteFarm(id int) (*models.Farm, error)
}

type farmService struct {
	farmRepository r.FarmRepository
}

type FConfig struct {
	FarmRepository r.FarmRepository
}

func NewFarmService(c *FConfig) FarmService {
	return &farmService{
		farmRepository: c.FarmRepository,
	}
}

func (s *farmService) GetFarm(query *dto.FarmRequestQuery) ([]*models.Farm, error) {
	farm, err := s.farmRepository.FindAll(query)
	if err != nil {
		return farm, err
	}

	return farm, nil
}

func (s *farmService) CountFarm() (int64, error) {
	totalFarms, err := s.farmRepository.Count()
	if err != nil {
		return totalFarms, err
	}

	return totalFarms, nil
}

func (s *farmService) CreateFarm(input *dto.FarmRequestBody) (*models.Farm, error) {
	farm := &models.Farm{}

	myFarm, err := s.farmRepository.FindByName(input.Name)
	if err != nil {
		return &models.Farm{}, err
	}

	if len(myFarm.Name) > 0 {
		return myFarm, &custom_error.FarmAlreadyExistsError{}
	}

	farm.Name = input.Name
	log.Println(farm)

	farm, err = s.farmRepository.Save(farm)
	if err != nil {
		return farm, err
	}

	return farm, nil
}

func (s *farmService) UpdateFarm(query *dto.FarmRequestBody, id int) (*models.Farm, error) {
	farm := &models.Farm{}

	myFarm, err := s.farmRepository.FindById(id)
	if err != nil {
		return &models.Farm{}, err
	}

	if myFarm.FarmId == 0 {
		return myFarm, &custom_error.FarmNotFoundError{}
	}

	myFarm, err = s.farmRepository.FindByName(query.Name)
	if err != nil {
		return &models.Farm{}, err
	}

	if myFarm.FarmId != 0 {
		return myFarm, &custom_error.FarmAlreadyExistsError{}
	}

	farm.Name = query.Name
	farm.FarmId = id
	farm, err = s.farmRepository.Update(farm)
	if err != nil {
		return farm, err
	}

	return farm, nil
}

func (s *farmService) DeleteFarm(id int) (*models.Farm, error) {
	farm := &models.Farm{}

	myFarm, err := s.farmRepository.FindById(id)
	if err != nil {
		return &models.Farm{}, err
	}

	if myFarm.FarmId == 0 {
		return myFarm, &custom_error.FarmNotFoundError{}
	}

	farm, err = s.farmRepository.Delete(id)
	if err != nil {
		return farm, err
	}

	return farm, nil
}
