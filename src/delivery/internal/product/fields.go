package product

import (
	"fmt"
	"regexp"

	"github.com/joseluis8906/go-code/src/pkg/types"
)

const (
	NameField = "Name"
)

// Name is a product name
type (
	Name types.StringValue
	Ref  types.StringValue
)

func NewName(value string) (Name, error) {
	if len(value) < 3 {
		return Name{}, fmt.Errorf("invalid product name")
	}

	return Name{value, true}, nil
}

func NewRef(value string) (Ref, error) {
	re := regexp.MustCompile(`^[A-Z]{3}\-[0-9]{3}$`)
	if !re.MatchString(value) {
		return Ref{}, fmt.Errorf("invalid product reference")
	}

	return Ref{value, true}, nil
}
