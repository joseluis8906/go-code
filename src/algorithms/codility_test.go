package algorithms_test

import (
	"testing"

	"github.com/joseluis8906/go-code/src/algorithms"
	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Parallel()

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
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := algorithms.Solution(tc.input)

			assert.Equal(t, tc.want, got)
		})
	}
}
