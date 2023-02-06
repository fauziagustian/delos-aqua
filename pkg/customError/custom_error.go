package custom_error

type UserNotFoundError struct {
}

type UserAlreadyExistsError struct {
}

type NotValidEmailError struct {
}

type IncorrectCredentialsError struct {
}

func (e *IncorrectCredentialsError) Error() string {
	return "incorrect email or password"
}

func (e *NotValidEmailError) Error() string {
	return "not a valid email"
}

func (e *UserNotFoundError) Error() string {
	return "user not found"
}
func (e *UserAlreadyExistsError) Error() string {
	return "user already exists"
}
