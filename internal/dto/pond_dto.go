package dto

type PondResponseBody struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Farm string `json:"farm_id"`
}
