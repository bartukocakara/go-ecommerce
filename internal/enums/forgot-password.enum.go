package enums

// ForgotPasswordResult represents the different resulting cases for forgot password.
type ForgotPasswordResult int

const (
	UserNotFound ForgotPasswordResult = iota
	TokenAlreadyExists
	TokenCouldntGenerated
	Success
	UnknownError
)
