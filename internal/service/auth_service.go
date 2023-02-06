package service

import (
	"net/mail"

	"github.com/fauziagustian/delos-aqua/internal/dto"
	"github.com/fauziagustian/delos-aqua/internal/models"
	r "github.com/fauziagustian/delos-aqua/internal/repository"
	custom_error "github.com/fauziagustian/delos-aqua/pkg/customError"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Attempt(input *dto.LoginRequestBody) (*models.User, error)
}

type authService struct {
	userRepository r.UserRepository
}

type ASConfig struct {
	UserRepository r.UserRepository
}

func NewAuthService(c *ASConfig) AuthService {
	return &authService{
		userRepository: c.UserRepository,
	}
}

func (s *authService) Attempt(input *dto.LoginRequestBody) (*models.User, error) {
	_, err := mail.ParseAddress(input.Email)
	if err != nil {
		return &models.User{}, &custom_error.NotValidEmailError{}
	}

	user, err := s.userRepository.FindByEmail(input.Email)
	if err != nil {
		return user, err
	}

	if user.UserId == 0 {
		return user, &custom_error.UserNotFoundError{}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return user, &custom_error.IncorrectCredentialsError{}
	}

	return user, nil
}
