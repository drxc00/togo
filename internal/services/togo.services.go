package services

import (
	"drxc00/togo/internal/models"
	"drxc00/togo/internal/utils"
)

func CreateTogo(r models.TogoRequest) (models.Togo, bool) {
	// Generate a new ID for the to-go order
	id := utils.GenerateID()

	// Create TOGO
	togo := models.Togo{
		ID:          id,
		Location:    r.Location,
		Description: r.Description,
		CreatedAt:   utils.GetCurrentTime(),
		UpdatedAt:   utils.GetCurrentTime(),
		TogoDate: models.TogoDate{
			ArrivalDate:   r.TogoDate.ArrivalDate,
			DepartureDate: r.TogoDate.DepartureDate,
		},
	}

	return togo, true
}
