package enrollment

import "github.com/labstack/echo/v4"

type EnrollmentHandler interface {
	CreateEnrollment(c echo.Context) error
	GetEnrollments(c echo.Context) error
	GetEnrollment(c echo.Context) error
	MarkStepAsComplete(c echo.Context) error
	SetEnrollmentStatusAsDone(c echo.Context) error
	DeleteEnrollment(c echo.Context) error
}
