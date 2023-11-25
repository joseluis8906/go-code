package algorithms_test

import (
	"testing"

	"github.com/joseluis8906/go-code/algorithms"
	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "Five",
			input: []int{1, 3, 6, 4, 1, 2},
			want:  5,
		},
		{
			name:  "Four",
			input: []int{1, 2, 3},
			want:  4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := algorithms.Solution(tc.input)

			assert.Equal(t, tc.want, got)
		})
	}
}
