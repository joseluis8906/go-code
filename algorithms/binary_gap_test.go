package algorithms_test

import (
	"testing"

	"github.com/joseluis8906/yummies/go-code/src/algorithms"
	"github.com/joseluis8906/yummies/go-code/src/pkg/cmp"
)

func TestBinaryGap(t *testing.T) {
	testCases := map[string]struct {
		in   int
		want int
	}{
		"One": {
			in:   2,
			want: 1,
		},
		"ThirtyTwo": {
			in:   32,
			want: 5,
		},
		"ThrityTwo": {
			in:   64,
			want: 6,
		},
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := algorithms.BinaryGap(tc.in)

			if got != tc.want {
				t.Errorf("BinaryGap(%d) = %d; want %d", tc.in, got, tc.want, cmp.Diff(tc.want, got))
			}
		})
	}
}
