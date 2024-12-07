package entities

type EnrolledStep struct {
	Base
	UserID          string
	EnrolledPlantID string
	StepID          string
	Step            Step
	MarkComplete    bool	
}
