package mailer

// AppError represents application-level errors.
type AppError string

func (e AppError) Error() string {
	return string(e)
}

// AuthError represents authentication-related errors.
type AuthError string

func (e AuthError) Error() string {
	return string(e)
}

// EmailError represents email-related errors.
type EmailError string

func (e EmailError) Error() string {
	return string(e)
}

// Application-level errors
const (
	ErrInvalidConfig AppError = "invalid app config"
)

// Email errors
const (
	ErrEmptyTo EmailError = "email recipient(s) is empty"
)
