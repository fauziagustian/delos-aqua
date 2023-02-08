package service

import (
	"github.com/fauziagustian/delos-aqua/internal/dto"
	"github.com/fauziagustian/delos-aqua/internal/models"
	r "github.com/fauziagustian/delos-aqua/internal/repository"
	custom_error "github.com/fauziagustian/delos-aqua/pkg/customError"
)

type PondService interface {
	GetPond(query *dto.RequestQuery) ([]*models.Pond, error)
	CountPond() (int64, error)
	CreatePond(query *dto.PondRequestBody) (*models.Pond, error)
	UpdatePond(query *dto.PondRequestBody, id int) (*models.Pond, error)
	DeletePond(id int) (*models.Pond, error)
}

type pondService struct {
	pondRepository r.PondRepository
}

type PConfig struct {
	PondRepository r.PondRepository
}

func NewPondService(c *PConfig) PondService {
	return &pondService{
		pondRepository: c.PondRepository,
	}
}

func (s *pondService) GetPond(query *dto.RequestQuery) ([]*models.Pond, error) {
	pond, err := s.pondRepository.FindAll(query)
	if err != nil {
		return pond, err
	}

	return pond, nil
}

func (s *pondService) CountPond() (int64, error) {
	totalPonds, err := s.pondRepository.Count()
	if err != nil {
		return totalPonds, err
	}

	return totalPonds, nil
}

func (s *pondService) CreatePond(input *dto.PondRequestBody) (*models.Pond, error) {
	pond := &models.Pond{}

	myPond, err := s.pondRepository.FindByName(input.Name)
	if err != nil {
		return &models.Pond{}, err
	}

	if len(myPond.Name) > 0 {
		return myPond, &custom_error.PondAlreadyExistsError{}
	}

	pond.Name = input.Name
	pond.FarmID = input.FarmId
	pond, err = s.pondRepository.Save(pond)
	if err != nil {
		return pond, err
	}

	return pond, nil
}

func (s *pondService) UpdatePond(query *dto.PondRequestBody, id int) (*models.Pond, error) {
	pond := &models.Pond{}

	myPondId, err := s.pondRepository.FindById(id)
	if err != nil {
		return &models.Pond{}, err
	}

	if myPondId.PondId == 0 {
		return myPondId, &custom_error.PondNotFoundError{}
	}

	myPondName, err := s.pondRepository.FindByName(query.Name)
	if err != nil {
		return &models.Pond{}, err
	}

	if myPondName.PondId != 0 && myPondId.Name != myPondName.Name {
		return myPondName, &custom_error.PondAlreadyExistsError{}
	}

	pond.Name = query.Name
	pond.PondId = id
	pond.FarmID = query.FarmId
	pond, err = s.pondRepository.Update(pond)
	if err != nil {
		return pond, err
	}

	return pond, nil
}

func (s *pondService) DeletePond(id int) (*models.Pond, error) {
	pond := &models.Pond{}

	myPond, err := s.pondRepository.FindById(id)
	if err != nil {
		return &models.Pond{}, err
	}

	if myPond.PondId == 0 {
		return myPond, &custom_error.PondNotFoundError{}
	}

	pond, err = s.pondRepository.Delete(id)
	if err != nil {
		return pond, err
	}

	return pond, nil
}
