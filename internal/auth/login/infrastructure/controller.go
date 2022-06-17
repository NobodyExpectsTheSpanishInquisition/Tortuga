package infrastructure

import (
	"Tortuga/internal/auth/login/application"
	. "Tortuga/internal/auth/shared/application"
)

func Login(request LoginRequest, validator PasswordValidatorInterface, emailValidator EmailValidatorInterface, authenticator AuthenticatorInterface) (string, error) {
	var err error
	email, password := request.Email, request.Password

	err = emailValidator.Validate(email)
	err = validator.Validate(password)

	query := application.NewLoginQuery(email, password)

	token, err := application.Handle(query, authenticator)
	if err != nil {
		return "", err
	}

	return token, err
}

type LoginRequest struct {
	Email    string
	Password string
}
