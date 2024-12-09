package enrollment

import "go-agreenery/entities"

type EnrollmentRepository interface {
	CreateEnrollment(plant entities.EnrolledPlant) (entities.EnrolledPlant, error)
	GetEnrollments(filter entities.Filter) ([]entities.EnrolledPlant, entities.Pagination, error)
	GetEnrollment(enrollmentID, currUserID string) (entities.EnrolledPlant, error)
	MarkStepAsComplete(stepID, currUserID string) (entities.EnrolledPlant, error)
	SetEnrollmentStatusAsDone(enrollmentID, currUserID string) (entities.EnrolledPlant, error)
	DeleteEnrollment(enrollmentID, currUserID string) error
}
