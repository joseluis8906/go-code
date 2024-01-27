package product

import (
	"fmt"
	"regexp"
)

const (
	NameField = "Name"
)

// Name is a product name
type Name struct {
	value string
}

func (n Name) String() string {
	return n.value
}

func NewName(value string) (Name, error) {
	if len(value) < 3 {
		return Name{value: "<nil>"}, fmt.Errorf("invalid product name")
	}

	return Name{value}, nil
}

type Ref struct {
	value string
}

func (r Ref) String() string {
	return r.value
}

func NewRef(value string) (Ref, error) {
	re := regexp.MustCompile(`^[A-Z][3]\-[0-9][3]$`)

	if !re.MatchString(value) {
		return Ref{value: "<nil>"}, fmt.Errorf("invalid product reference")
	}

	return Ref{value}, nil
}
