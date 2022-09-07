package services

type ErrUsernameNotUnique struct{}

func (ErrUsernameNotUnique) Error() string {
	return "User with the specified username already exists"
}

type ErrEmailNotUnique struct{}

func (ErrEmailNotUnique) Error() string {
	return "User with the specified email already exists"
}
