package dto

import (
	"github.com/fauziagustian/delos-aqua/internal/models"
)

type LoginRequestBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5"`
}

type RegisterRequestBody struct {
	Name     string `json:"name" binding:"required,alphanum"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5"`
}

type ForgotPasswordRequestBody struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordRequestBody struct {
	Token           string `json:"token" binding:"required"`
	Password        string `json:"password" binding:"required,min=5"`
	ConfirmPassword string `json:"confirm_password" binding:"required,min=5"`
}

type ForgotPasswordResponseBody struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type LoginResponseBody struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FormatLogin(user *models.User, token string) LoginResponseBody {
	return LoginResponseBody{
		ID:    uint(user.UserId),
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}
}
