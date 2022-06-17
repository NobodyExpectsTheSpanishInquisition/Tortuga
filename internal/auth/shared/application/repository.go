package application

import Entity "Tortuga/internal/Infrastructure"

type UserAuthRepositoryInterface interface {
	FindByEmail(email string) (Entity.User, error)
}
