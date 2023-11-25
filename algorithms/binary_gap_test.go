package algorithms_test

import (
	"testing"

	"github.com/joseluis8906/go-code/algorithms"
	"github.com/stretchr/testify/assert"
)

func TestBinaryGap(t *testing.T) {
	testCases := []struct {
		name  string
		input int
		want  int
	}{
		{
			name:  "One",
			input: 2,
			want:  1,
		},
		{
			name:  "ThirtyTwo",
			input: 32,
			want:  5,
		},
		{
			name:  "ThrityTwo",
			input: 64,
			want:  6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := algorithms.BinaryGap(tc.input)

			assert.Equal(t, tc.want, got)
		})
	}
}
