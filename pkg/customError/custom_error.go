package custom_error

type UserNotFoundError struct {
}

type UserAlreadyExistsError struct {
}

type NotValidEmailError struct {
}

type IncorrectCredentialsError struct {
}

type FarmAlreadyExistsError struct {
}

type FarmNotFoundError struct {
}

type PondAlreadyExistsError struct {
}

type PondNotFoundError struct {
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

func (e *FarmAlreadyExistsError) Error() string {
	return "farm already exists"
}

func (e *FarmNotFoundError) Error() string {
	return "farm not found"
}

func (e *PondAlreadyExistsError) Error() string {
	return "pond already exists"
}

func (e *PondNotFoundError) Error() string {
	return "pond not found"
}
