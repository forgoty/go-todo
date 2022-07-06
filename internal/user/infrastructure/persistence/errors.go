package persistence

type persistenceError string

func (e persistenceError) Error() string {
	return string(e)
}

const (
	ErrInvalidCredsOrNotFound persistenceError = "No account found with the given credentials"
	UserNotFound              persistenceError = "User Not Found"
)
