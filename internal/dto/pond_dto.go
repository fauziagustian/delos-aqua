package dto

import (
	"time"

	"github.com/fauziagustian/delos-aqua/internal/models"
)

type PondResponseBody struct {
	PondId uint   `json:"pond_id"`
	Name   string `json:"name"`
	FarmId string `json:"farm_id"`
}

type PondRequestBody struct {
	Name   string `json:"name"`
	FarmId int    `json:"farm_id"`
}

type PondResponse struct {
	PondId    uint      `json:"pond_id"`
	Name      string    `json:"farm"`
	FarmId    uint      `json:"farm_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatPond(pond *models.Pond) PondResponse {
	return PondResponse{
		PondId:    uint(pond.PondId),
		Name:      pond.Name,
		FarmId:    uint(pond.FarmID),
		CreatedAt: pond.CreatedAt,
		UpdatedAt: pond.UpdatedAt,
	}
}

func FormatPonds(ponds []*models.Pond) []PondResponse {
	formattedPond := []PondResponse{}
	for _, pond := range ponds {
		formattedBook := FormatPond(pond)
		formattedPond = append(formattedPond, formattedBook)
	}
	return formattedPond
}
