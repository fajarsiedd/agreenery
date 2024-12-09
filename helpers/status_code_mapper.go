package helpers

import (
	c "go-agreenery/constants"
	"net/http"
	"strings"
)

func GetStatusCodeBySuccessMessage(msg string) int {
	if strings.Contains(msg, "registered") || strings.Contains(msg, "created") {
		return http.StatusCreated
	} else {
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
