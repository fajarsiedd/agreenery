package enrollment

import (
	"go-agreenery/entities"
	"go-agreenery/repositories/enrollment"
)

type enrollmentService struct {
	repository enrollment.EnrollmentRepository
}

func NewEnrollmentService(r enrollment.EnrollmentRepository) *enrollmentService {
	return &enrollmentService{
		repository: r,
	}
}

func (s enrollmentService) CreateEnrollment(plant entities.EnrolledPlant) (entities.EnrolledPlant, error) {
	return s.repository.CreateEnrollment(plant)
}

func (s enrollmentService) GetEnrollments(filter entities.Filter) ([]entities.EnrolledPlant, entities.Pagination, error) {
	return s.repository.GetEnrollments(filter)
}

func (s enrollmentService) GetEnrollment(enrollmentID string) (entities.EnrolledPlant, error) {
	return s.repository.GetEnrollment(enrollmentID)
}

func (s enrollmentService) MarkStepAsComplete(stepID string) (entities.EnrolledPlant, error) {
	return s.repository.MarkStepAsComplete(stepID)
}

func (s enrollmentService) SetEnrollmentStatusAsDone(enrollmentID string) (entities.EnrolledPlant, error) {
	return s.repository.SetEnrollmentStatusAsDone(enrollmentID)
}

func (s enrollmentService) DeleteEnrollment(enrollmentID string) error {
	return s.repository.DeleteEnrollment(enrollmentID)
}
