package application

import (
	model "Tortuga/internal/auth/shared/domain"
)

type AuthenticatorInterface interface {
	Authenticate(email string, password string) (model.AuthenticatedUser, error)
	CreateToken(user model.AuthenticatedUser) (string, error)
}
