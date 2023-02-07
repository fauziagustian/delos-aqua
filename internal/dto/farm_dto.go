package dto

import (
	"strings"
	"time"

	"github.com/fauziagustian/delos-aqua/internal/models"
	"gorm.io/gorm"
)

type FarmRequestBody struct {
	Name string `json:"name"`
}

type FarmResponseBody struct {
	FarmId uint   `json:"farm_id"`
	Name   string `json:"name"`
}

type FarmRequestQuery struct {
	Search string `form:"s"`
	SortBy string `form:"sortBy"`
	Sort   string `form:"sort"`
	Limit  int    `form:"limit"`
	Page   int    `form:"page"`
}

type FarmResponse struct {
	FarmId    uint      `json:"id"`
	Name      string    `json:"farm"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FarmResponseDelete struct {
	FarmId    uint           `json:"id"`
	Name      string         `json:"farm"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func FormatQuery(query *FarmRequestQuery) *FarmRequestQuery {
	if query.Limit == 0 {
		query.Limit = 10
	}
	if query.Page == 0 {
		query.Page = 1
	}

	query.SortBy = strings.ToLower(query.SortBy)
	if query.SortBy == "date" {
		query.SortBy = "updated_at"
	} else if query.SortBy == "to" {
		query.SortBy = "destination_id"
	} else if query.SortBy != "amount" {
		query.SortBy = "updated_at"
	}

	query.Sort = strings.ToUpper(query.Sort)
	if query.Sort != "ASC" {
		query.Sort = "DESC"
	}

	return query
}

func FormatFarm(farm *models.Farm) FarmResponse {
	return FarmResponse{
		FarmId:    uint(farm.FarmId),
		Name:      farm.Name,
		CreatedAt: farm.CreatedAt,
		UpdatedAt: farm.UpdatedAt,
	}
}

func FormatDeleteFarm(farm *models.Farm) FarmResponseDelete {
	return FarmResponseDelete{
		FarmId:    uint(farm.FarmId),
		Name:      farm.Name,
		CreatedAt: farm.CreatedAt,
		UpdatedAt: farm.UpdatedAt,
		DeletedAt: farm.DeletedAt,
	}
}

func FormatFarms(transactions []*models.Farm) []FarmResponse {
	formattedFarm := []FarmResponse{}
	for _, transaction := range transactions {
		formattedBook := FormatFarm(transaction)
		formattedFarm = append(formattedFarm, formattedBook)
	}
	return formattedFarm
}
