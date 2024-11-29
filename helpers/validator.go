package helpers

import (
	"errors"
	"go-agreenery/constants"
	"strings"

	"github.com/go-playground/validator/v10"
)

func TranslateValidationErr(err error) error {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		msg := strings.ToLower(ve[0].Field() + " " + msgForTag(ve[0].Tag()))

		return errors.New(msg)
	}

	return nil
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return constants.ErrFieldRequired.Error()
	case "email":
		return constants.ErrInvalidFormat.Error()
	}
	return ""
}
