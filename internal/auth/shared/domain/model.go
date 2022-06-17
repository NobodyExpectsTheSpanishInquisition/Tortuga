package domain

type AuthenticatedUser struct {
	id       string
	email    string
	password string
}

func NewAuthenticatedUser(email string, password string) *AuthenticatedUser {
	return &AuthenticatedUser{email: email, password: password}
}

func (a AuthenticatedUser) Id() string {
	return a.id
}

func (a AuthenticatedUser) Email() string {
	return a.email
}

func (a AuthenticatedUser) Password() string {
	return a.password
}
