package application

type PasswordValidatorInterface interface {
	Validate(password string) error
}

type EmailValidatorInterface interface {
	Validate(email string) error
}

type InvalidPasswordError struct {
}

func (err InvalidPasswordError) Error() string {
	return err.Error()
}
