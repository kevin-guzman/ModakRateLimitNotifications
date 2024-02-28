package model

type AbuseError interface {
	error
	Details() string
}
type applicationAbuseError struct {
	err     error
	details string
}

func NewAbuseError(err error, details string) AbuseError {
	return applicationAbuseError{
		err,
		details,
	}
}
func (e applicationAbuseError) Details() string {
	return e.details
}
func (e applicationAbuseError) Error() string {
	return e.err.Error()
}
