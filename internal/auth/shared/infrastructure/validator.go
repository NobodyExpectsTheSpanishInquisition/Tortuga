package infrastructure

import (
	"github.com/go-passwd/validator"
	"net/mail"
)

type PasswordValidator struct {
	validator validator.Validator
}

type EmailValidator struct {
}

func NewEmailValidator() EmailValidator {
	return EmailValidator{}
}

func (e EmailValidator) Validate(email string) error {
	_, err := mail.ParseAddress(email)

	return err
}

func NewPasswordValidator(validator validator.Validator) PasswordValidator {
	return PasswordValidator{validator: validator}
}

func (v PasswordValidator) Validate(password string) error {
	return v.validator.Validate(password)
}
