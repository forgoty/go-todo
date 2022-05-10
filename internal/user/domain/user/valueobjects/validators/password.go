package validators

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidatePassword(password string) error {
	return validation.Validate(
		password,
		validation.Required,
		validation.Length(8, 125),
	)
}
