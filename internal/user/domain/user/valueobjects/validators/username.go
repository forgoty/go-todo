package validators

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateUsername(username string) error {
	return validation.Validate(
		username,
		validation.Required,
		validation.Length(5, 125),
	)
}
