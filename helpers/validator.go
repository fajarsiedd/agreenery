package helpers

import (
	"errors"
	"go-agreenery/constants"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

func TranslateValidationErr(err error) error {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		msg := toSnakeCase(ve[0].Field() + " " + msgForTag(ve[0].Tag()))

		return errors.New(msg)
	}

	return nil
}

func toSnakeCase(str string) string {
	if len(str) == 1 || !regexp.MustCompile("[A-Z]").MatchString(str) {
		return strings.ToLower(str)
	}

	re := regexp.MustCompile("(^|[a-z])([A-Z])")

	str = re.ReplaceAllString(str, "${1}_${2}")

	return strings.ToLower(str[1:])
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return constants.ErrFieldRequired.Error()
	case "email":
		return constants.ErrInvalidFormat.Error()
	}
	return constants.ErrFieldInvalid.Error()
}
