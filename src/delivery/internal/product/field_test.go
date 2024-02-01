package product

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestName(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		t.Parallel()

		in := "Hmaburger"
		want := Name{in, true}

		got, err := NewName(in)

		if got != want || err != nil {
			diff := fmt.Sprintf("diff:\n\t-want\n\t+got\n%v", cmp.Diff(want, got))
			t.Errorf("NewName(%v) = %v, %v; want %v, nil\n%v", in, got, err, want, diff)
		}
	})

	t.Run("TooShort", func(t *testing.T) {
		t.Parallel()

		in := "Ha"
		var want Name

		got, err := NewName(in)

		if got != want || err == nil {
			diff := fmt.Sprintf("diff:\n\t-want\n\t+got\n%v", cmp.Diff(want, got))
			t.Errorf("NewName(%v) = %v, %v; want %v, error\n%v", in, got, err, want, diff)
		}
	})
}

func TestReference(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		t.Parallel()

		in := "HAM-123"
		want := Ref{in, true}

		got, err := NewRef(in)

		if got != want || err != nil {
			diff := fmt.Sprintf("diff:\n\t-want\n\t+got\n%v", cmp.Diff(want, got))
			t.Errorf("NewRef(%v) = %v, %v; want %v, nil\n%v", in, got, err, want, diff)
		}
	})

	t.Run("Invalid", func(t *testing.T) {
		var zeroRef Ref

		testCases := map[string]struct {
			in   string
			want Ref
		}{
			"TooShortLetters": {
				in:   "AB-123",
				want: zeroRef,
			},
			"TooShortNumbers": {
				in:   "ABC-12",
				want: zeroRef,
			},
		}

		for name, tc := range testCases {
			tc := tc

			t.Run(name, func(t *testing.T) {
				t.Parallel()

				got, err := NewRef(tc.in)

				if got != tc.want || err == nil {
					diff := fmt.Sprintf("diff:\n\t-want\n\t+got\n%v", cmp.Diff(tc.want, got))
					t.Errorf("NewRef(%v) = %v, %v; want %v, error\n%v", tc.in, got, err, tc.want, diff)
				}
			})
		}
	})
}
