package customer

import (
	"errors"
	"regexp"
)

const (
	EmailField = "Email"
)

// Email represents a customer's email address.
type Email struct {
	value string
}

func (e Email) String() string {
	return e.value
}

// NewEmail creates a new Email.
func NewEmail(value string) (Email, error) {
	re := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	if !re.MatchString(value) {
		return Email{value: "<nil>"}, errors.New("invalid email address")
	}

	return Email{value}, nil
}
