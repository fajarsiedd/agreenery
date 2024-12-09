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

func (s enrollmentService) GetEnrollment(enrollmentID, currUserID string) (entities.EnrolledPlant, error) {
	return s.repository.GetEnrollment(enrollmentID, currUserID)
}

func (s enrollmentService) MarkStepAsComplete(stepID, currUserID string) (entities.EnrolledPlant, error) {
	return s.repository.MarkStepAsComplete(stepID, currUserID)
}

func (s enrollmentService) SetEnrollmentStatusAsDone(enrollmentID, currUserID string) (entities.EnrolledPlant, error) {
	return s.repository.SetEnrollmentStatusAsDone(enrollmentID, currUserID)
}

func (s enrollmentService) DeleteEnrollment(enrollmentID, currUserID string) error {
	return s.repository.DeleteEnrollment(enrollmentID, currUserID)
}
