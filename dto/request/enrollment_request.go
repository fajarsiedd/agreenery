package request

import "go-agreenery/entities"

type EnrollmentRequest struct {
	PlantID string `json:"plant_id" validate:"required"`
	UserID  string
}

func (r EnrollmentRequest) ToEntity() entities.EnrolledPlant {
	return entities.EnrolledPlant{
		UserID: r.UserID,
		PlantID: r.PlantID,
	}
}
