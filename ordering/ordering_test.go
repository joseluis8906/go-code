package ordering_test

import (
	"testing"

	"github.com/joseluis8906/yummies/go-code/src/pkg/ordering"
)

func TestOrdering(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		left  ordering.Cmp
		right ordering.Cmp
	}{
		"Less": {
			left:  ordering.Less,
			right: ordering.Equal,
		},
		"Equal": {
			left:  ordering.Equal,
			right: ordering.Greater,
		},
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if !(tc.left < tc.right) {
				t.Errorf("left should be less than right")
			}
		})
	}
}
