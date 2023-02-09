package dto

import (
	"strings"
	"time"

	"github.com/fauziagustian/delos-aqua/internal/models"
)

type FarmRequestBody struct {
	Name string `json:"name"`
}

type FarmResponseBody struct {
	FarmId uint   `json:"farm_id"`
	Name   string `json:"name"`
}

type RequestQuery struct {
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

func FormatQuery(query *RequestQuery) *RequestQuery {
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

func FormatFarms(farms []*models.Farm) []FarmResponse {
	formattedFarm := []FarmResponse{}
	for _, farm := range farms {
		formattedBook := FormatFarm(farm)
		formattedFarm = append(formattedFarm, formattedBook)
	}
	return formattedFarm
}
