package ncispec

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func init() {
	validate = validator.New()
	_ = validate.RegisterValidation("is-slug", validateSlug)
	_ = validate.RegisterValidation("is-arch", validateArch)
}

type validationError struct {
	Field       string
	Value       string
	Description string
}

func validateSlug(fl validator.FieldLevel) bool {
	return slugRegex.MatchString(fl.Field().String())
}

func validateArch(fl validator.FieldLevel) bool {
	return archRegex.MatchString(fl.Field().String())
}

func validateSpec(spec *NormalizeCISpec) []validationError {
	var errors []validationError

	err := validate.Struct(spec)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Err(err).Msg("invalid value for validation")
			return errors
		}

		validationErrors := err.(validator.ValidationErrors)
		for _, err := range validationErrors {
			errors = append(errors, validationError{err.Field(), fmt.Sprintf("%v", err.Value()), err.Tag()})
		}
	}

	return errors
}
