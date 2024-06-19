package algorithms_test

import (
	"testing"

	"github.com/joseluis8906/yummies/go-code/src/algorithms"
)

func TestLongestPalindromicSubstr(t *testing.T) {
	testCases := map[string]struct {
		in   string
		want string
	}{
		"geeksskeeg": {
			in:   "forgeeksskeegfor",
			want: "geeksskeeg",
		},
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			got := algorithms.LongestPalindromicSubstr(tc.in)

			if got != tc.want {
				t.Errorf("LongestPalindromicSubstr(%s) = %s; want %s", tc.in, got, tc.want)
			}
		})
	}
}
