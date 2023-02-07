package dto

import "github.com/fauziagustian/delos-aqua/internal/models"

type UserRequestParams struct {
	UserID int `uri:"id" binding:"required"`
}

type UserRequestQuery struct {
	Name  string `form:"name"`
	Email string `form:"email"`
}

type UserResponseBody struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserDetailResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FormatUserDetail(user *models.User) UserDetailResponse {
	formattedUser := UserDetailResponse{}
	formattedUser.ID = uint(user.UserId)
	formattedUser.Name = user.Name
	formattedUser.Email = user.Email
	return formattedUser
}
