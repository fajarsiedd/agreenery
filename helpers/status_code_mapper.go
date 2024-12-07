package helpers

import (
	c "go-agreenery/constants"
	"net/http"
	"strings"
)

func GetStatusCodeBySuccessMessage(msg string) int {
	switch msg {
	case c.RegisterSuccess, c.CreateCategorySuccess, c.CreatePlantSuccess, c.CreateStepSuccess, c.CreateEnrollmentSuccess:
		return http.StatusCreated
	default:
		return http.StatusOK
	}
}

func GetStatusCodeByErr(err error) int {
	if errMsg := err.Error(); strings.Contains(errMsg, "required") || strings.Contains(errMsg, "invalid") {
		return http.StatusBadRequest
	} else if err == c.ErrAccessNotAllowed {
		return http.StatusUnauthorized
	} else {
		return http.StatusInternalServerError
	}
}
