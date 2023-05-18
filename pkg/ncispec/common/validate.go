package common

import (
	"github.com/go-playground/validator/v10"
)

// Validator is the validator instance used for validating the spec
var Validator *validator.Validate

func init() {
	Validator = validator.New()
	_ = Validator.RegisterValidation("is-slug", validateSlug)
	_ = Validator.RegisterValidation("is-arch", validateArch)
}

type ValidationError struct {
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
