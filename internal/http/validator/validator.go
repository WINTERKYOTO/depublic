// package validator

package validator

import (
	"depublic/entity"

	"github.com/go-playground/validator/v10"
)

// ValidateUser validates a `User`
func ValidateUser(user *entity.User) error {
	// Create the validator
	v := validator.New()

	// Validate the user
	err := v.Struct(user)
	if err != nil {
		return err
	}

	return nil
}
