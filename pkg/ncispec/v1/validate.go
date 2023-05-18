package v1

import (
	"fmt"

	"github.com/cidverse/normalizeci/pkg/ncispec/common"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// Validate will validate the spec for completion
func (spec Spec) Validate() []common.ValidationError {
	var errors []common.ValidationError

	err := common.Validator.Struct(spec)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Err(err).Msg("invalid value for validation")
			return errors
		}

		validationErrors := err.(validator.ValidationErrors)
		for _, err := range validationErrors {
			errors = append(errors, common.ValidationError{Field: err.Field(), Value: fmt.Sprintf("%v", err.Value()), Description: err.Tag()})
		}
	}

	return errors
}
