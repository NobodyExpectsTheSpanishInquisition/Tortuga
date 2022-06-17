package Infrastructure

import (
	auth "Tortuga/internal/auth/shared/application"
	"Tortuga/internal/auth/shared/infrastructure"
	externalPasswordValidator "github.com/go-passwd/validator"
)

type Container struct {
	emailValidator    auth.EmailValidatorInterface
	passwordValidator auth.PasswordValidatorInterface
	authenticator     auth.AuthenticatorInterface
}

func (c Container) EmailValidator() auth.EmailValidatorInterface {
	return c.emailValidator
}

func (c Container) PasswordValidator() auth.PasswordValidatorInterface {
	return c.passwordValidator
}

func (c Container) Init() {
	c.emailValidator = buildEmailValidator()
	c.passwordValidator = buildPasswordValidator()
	c.authenticator = buildAuthenticator()
}

func buildAuthenticator() auth.AuthenticatorInterface {
	return
}

func buildPasswordValidator() infrastructure.PasswordValidator {
	err := auth.InvalidPasswordError{}
	return infrastructure.NewPasswordValidator(*externalPasswordValidator.New(externalPasswordValidator.MinLength(8, err)))
}

func buildEmailValidator() infrastructure.EmailValidator {
	return infrastructure.NewEmailValidator()
}
