package validator

import (
	"gopkg.in/go-playground/validator.v10"
)

// ValidateEmail validates an email address.
func ValidateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	if len(email) == 0 {
		return true
	}

	return validator.IsEmail(email)
}

// ValidatePassword validates a password.
func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if len(password) == 0 {
		return true
	}

	// Check if the password is at least 8 characters long
	if len(password) < 8 {
		return false
	}

	// Check if the password has at least one lowercase letter
	hasLowercase := false
	for _, c := range password {
		if c >= 'a' && c <= 'z' {
			hasLowercase = true
			break
		}
	}
	if !hasLowercase {
		return false
	}

	// Check if the password has at least one uppercase letter
	hasUppercase := false
	for _, c := range password {
		if c >= 'A' && c <= 'Z' {
			hasUppercase = true
			break
		}
	}
	if !hasUppercase {
		return false
	}

	// Check if the password has at least one number
	hasNumber := false
	for _, c := range password {
		if c >= '0' && c <= '9' {
			hasNumber = true
			break
		}
	}
	if !hasNumber {
		return false
	}

	// Check if the password has at least one special character
	hasSpecial := false
	for _, c := range password {
		if (c >= '!' && c <= '~') || c == '-' || c == '_' {
			hasSpecial = true
			break
		}
	}
	if !hasSpecial {
		return false
	}

	return true
}
