package financial_test

import (
	"testing"

	"github.com/joseluis8906/go-code/src/pkg/cmp"
	"github.com/joseluis8906/go-code/src/pkg/financial"
)

func TestMoney_Multiplication(t *testing.T) {
	five := financial.Dollar(5)

	testCases := map[string]struct {
		in   int
		want financial.Money
	}{
		"5*2": {in: 2, want: financial.Dollar(10)},
		"5*3": {in: 3, want: financial.Dollar(15)},
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := five.Times(tc.in)
			if got != tc.want {
				t.Errorf("[5 Dollar]money.Times(%d) = %v; want %v\n%v", tc.in, got, tc.want, cmp.Diff(tc.want))
			}
		})
	}
}

func TestMoney_Equality(t *testing.T) {
	testcases := map[string]struct {
		in   financial.Money
		want bool
	}{
		"5 Dollar == 5 Dollar": {
			in:   financial.Dollar(5),
			want: true,
		},
		"5 Dollar != 6 Dollar": {
			in:   financial.Dollar(6),
			want: false,
		},
		"5 Franc == 5 Franc": {
			in:   financial.Franc(5),
			want: true,
		},
	}

	for name, tc := range testcases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := tc.in.Equals(tc.in)

			if got != tc.want {
				t.Errorf("money.Equals(%v) = %v; want %v\n%v", tc.in, got, tc.want, cmp.Diff(tc.want, got))
			}
		})
	}
}

func TestMoney_Plus(t *testing.T) {
	five := financial.Dollar(5)

	testCases := map[string]struct {
		in   financial.Money
		want financial.Money
	}{
		"5 Dollar + 5 Dollar": {
			in:   financial.Dollar(5),
			want: financial.Dollar(10),
		},
		"5 Dollar + 5 Franc": {
			in:   financial.Franc(5),
			want: financial.Dollar(0),
		},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := five.Plus(tc.in)

			if got != tc.want {
				t.Errorf("[5 Dollar]money.Plus(%v) = %v; want %v\n%v", tc.in, got, tc.want, cmp.Diff(tc.want, got))
			}
		})
	}
}

func TestMoney_ValidCurrencies(t *testing.T) {
	testCases := map[string]struct {
		amount   int
		currency financial.Currency
		want     financial.Money
	}{
		"5 Dollar": {
			amount:   5,
			currency: financial.USD,
			want:     financial.Dollar(5),
		},
		"Zero": {
			amount:   5,
			currency: financial.Currency("UNKNOWN"),
			want:     financial.Money{},
		},
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := financial.NewMoney(tc.amount, tc.currency)

			if got != tc.want {
				t.Errorf("money.NewMoney(%d, %s) = %v; want %v\n%v", tc.amount, tc.currency, got, tc.want, cmp.Diff(tc.want, got))
			}
		})
	}
}
