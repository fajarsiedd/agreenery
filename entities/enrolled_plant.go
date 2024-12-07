package entities

type EnrolledPlant struct {
	Base
	UserID        string
	PlantID       string
	Plant         Plant	
	EnrolledSteps []EnrolledStep
	IsDone        bool
}
