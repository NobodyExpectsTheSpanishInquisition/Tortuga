package application

import authenticator "Tortuga/internal/auth/shared/application"

func Handle(query LoginQuery, authenticator authenticator.AuthenticatorInterface) (string, error) {
	var err error

	user, err := authenticator.Authenticate(query.email, query.password)

	if nil != err {
		return "", err
	}

	return authenticator.CreateToken(user)
}

type LoginQuery struct {
	email    string
	password string
}

func NewLoginQuery(email string, password string) LoginQuery {
	return LoginQuery{email: email, password: password}
}
