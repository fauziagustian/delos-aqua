package service

import (
	"github.com/fauziagustian/delos-aqua/internal/dto"
	r "github.com/fauziagustian/delos-aqua/internal/repository"
)

type UserAgentService interface {
	GetUserAgent() ([]*dto.UserAgentResponseCount, error)
}

type userAgentService struct {
	userAgentRepository r.UserAgentRepository
}

type UAConfig struct {
	UserAgentRepository r.UserAgentRepository
}

func NewUserAgentService(c *UAConfig) UserAgentService {
	return &userAgentService{
		userAgentRepository: c.UserAgentRepository,
	}
}

func (s *userAgentService) GetUserAgent() ([]*dto.UserAgentResponseCount, error) {
	userAgent, err := s.userAgentRepository.FindAll()
	if err != nil {
		return userAgent, err
	}

	return userAgent, nil
}
