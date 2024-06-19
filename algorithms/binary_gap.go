package algorithms

import (
	"fmt"
	"strings"
)

func BinaryGap(i int) int {
	x := fmt.Sprintf("%b", i)

	return strings.Count(x, "0")
}
