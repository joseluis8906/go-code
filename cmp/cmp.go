package cmp

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

var (
	diff  = cmp.Diff
	equal = cmp.Equal
)

func Diff(want, got interface{}) string {
	return fmt.Sprintf("diff:\n\t-want\n\t+got\n%v", cmp.Diff(want, got))
}

func Equal(want, got interface{}) bool {
	return equal(want, got)
}
