package models

import "encoding/json"

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type TogoDate struct {
	ArrivalDate   Date `json:"arrival_date"`
	DepartureDate Date `json:"departure_date"`
}

type Togo struct {
	ID          string   `json:"id"`
	Location    string   `json:"location"`
	Description string   `json:"description"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
	TogoDate    TogoDate `json:"togo_date"`
}

func (r *Response) Print() (string, error) {
	prettyJSON, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "", err
	}
	return string(prettyJSON), nil
}
