package models

type TogoRequest struct {
	Description string   `json:"description"`
	Location    string   `json:"location"`
	TogoDate    TogoDate `json:"togo_date"`
}

type Response struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
	Data    Togo   `json:"data"`
}
