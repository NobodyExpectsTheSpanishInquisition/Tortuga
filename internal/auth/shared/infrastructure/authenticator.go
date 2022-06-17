package infrastructure

import (
	"Tortuga/internal/auth/shared/application"
	model "Tortuga/internal/auth/shared/domain"
	"github.com/cristalhq/jwt/v4"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type JwtAuthenticator struct {
	userAuthRepository application.UserAuthRepositoryInterface
}

func (a JwtAuthenticator) createToken(user model.AuthenticatedUser) (string, error) {
	key := viper.GetString("JWT_SECRET")
	signer, err := jwt.NewSignerHS(jwt.HS256, []byte(key))

	if nil != err {
		return "", err
	}

	loginTime := time.Time{}
	claims := &jwt.RegisteredClaims{
		ID:        user.Id(),
		Issuer:    user.Email(),
		ExpiresAt: jwt.NewNumericDate(loginTime.Local()),
	}

	builder := jwt.NewBuilder(signer)

	token, err := builder.Build(claims)
	if nil != err {
		return "", err
	}

	return token.String(), err
}

func (a JwtAuthenticator) Authenticate(email string, password string) (model.AuthenticatedUser, error) {
	var err error
	var authenticatedUser model.AuthenticatedUser

	user, err := a.userAuthRepository.FindByEmail(email)
	if nil != err {
		return authenticatedUser, err
	}

	hashedPassword := user.Password
	err = a.checkPassword(password, hashedPassword)

	if nil != err {
		return authenticatedUser, err
	}

	return *model.NewAuthenticatedUser(user.Email, user.Password), err
}

func (a JwtAuthenticator) checkPassword(plainPassword string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}
