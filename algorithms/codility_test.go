package algorithms_test

import (
	"testing"

	"github.com/joseluis8906/yummies/go-code/src/algorithms"
	"github.com/joseluis8906/yummies/go-code/src/pkg/cmp"
)

func TestSolution(t *testing.T) {
	testCases := map[string]struct {
		in   []int
		want int
	}{
		"Five": {
			in:   []int{1, 3, 6, 4, 1, 2},
			want: 5,
		},
		"Four": {
			in:   []int{1, 2, 3},
			want: 4,
		},
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := algorithms.Solution(tc.in)

			if tc.want != got {
				t.Errorf("Solution(%v) = %v; want %v\n%v", tc.in, got, tc.want, cmp.Diff(tc.want, got))
			}
		})
	}
}
