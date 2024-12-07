package enrollment

import "go-agreenery/entities"

type EnrollmentService interface {
	CreateEnrollment(plant entities.EnrolledPlant) (entities.EnrolledPlant, error)
	GetEnrollments(filter entities.Filter) ([]entities.EnrolledPlant, entities.Pagination, error)
	GetEnrollment(enrollmentID string) (entities.EnrolledPlant, error)
	MarkStepAsComplete(stepID string) (entities.EnrolledPlant, error)
	SetEnrollmentStatusAsDone(enrollmentID string) (entities.EnrolledPlant, error)
	DeleteEnrollment(enrollmentID string) error
}
