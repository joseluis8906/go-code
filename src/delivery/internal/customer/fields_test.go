package customer

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEmail(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		t.Parallel()

		in := "john.doe@example.com"
		want := Email{"john.doe@example.com", true}

		got, err := NewEmail(in)
		if err != nil || want != got {
			diff := fmt.Sprintf("diff:\n\t-want\n\t+got\n%v", cmp.Diff(want, got))
			t.Errorf("NewEmail(%v) = %+v, %v; want %+v, nil\n%v", in, got, err, want, diff)
		}
	})

	t.Run("Invalid", func(t *testing.T) {
		var zeroEmail Email

		testCases := map[string]struct {
			in   string
			want Email
		}{
			"MissingDomain": {
				in:   "john.doe",
				want: zeroEmail,
			},

			"MissingName": {
				in:   "@example.com",
				want: zeroEmail,
			},

			"NotAllowedChars": {
				in:   "john#doe@example.com",
				want: zeroEmail,
			},
		}

		for name, tc := range testCases {
			tc := tc
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				got, err := NewEmail(tc.in)
				if err == nil || got != tc.want {
					diff := fmt.Sprintf("diff:\n\t-want\n\t+got\n%v", cmp.Diff(tc.want, got))
					t.Errorf("NewEmail(%v) = %v, %v; want %v, %v\n%v", tc.in, got, err, tc.want, "error", diff)
				}
			})
		}
	})
}
