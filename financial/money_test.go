package financial_test

import (
	"testing"

	"github.com/joseluis8906/yummies/go-code/src/pkg/cmp"
	"github.com/joseluis8906/yummies/go-code/src/pkg/financial"
)

func TestMoney_Multiplication(t *testing.T) {
	five := financial.Dollar(5)

	testCases := map[string]struct {
		in   int64
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
				t.Errorf("[5 Dollar]money.Times(%d) = %v; want %v\n%v", tc.in, got, tc.want, cmp.Diff(tc.want, got))
			}
		})
	}
}

func TestMoney_Equality(t *testing.T) {
	testcases := map[string]struct {
		left  financial.Money
		right financial.Money
		want  bool
	}{
		"5 Dollar == 5 Dollar": {
			left:  financial.Dollar(5),
			right: financial.Dollar(5),
			want:  true,
		},
		"5 Dollar != 6 Dollar": {
			left:  financial.Dollar(5),
			right: financial.Dollar(6),
			want:  false,
		},
		"5 Franc == 5 Franc": {
			left:  financial.Franc(5),
			right: financial.Franc(5),
			want:  true,
		},
	}

	for name, tc := range testcases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := tc.left.Equals(tc.right)

			if got != tc.want {
				t.Errorf("money.Equals(%v) = %v; want %v\n%v", tc.left, got, tc.want, cmp.Diff(tc.want, got))
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
			want: financial.Money{},
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
	t.Run("Valid currencies", func(t *testing.T) {
		amount := int64(5)
		currency := financial.USD

		want := financial.Dollar(amount)

		got, err := financial.NewMoney(amount, currency)

		if err != nil || got != want {
			t.Errorf("money.NewMoney(%d, %s) = %v %v; want %v nil\n%v", amount, currency, got, err, want, cmp.Diff(want, got))
		}
	})

	t.Run("Invalid currencies", func(t *testing.T) {
		amount := int64(5)
		currency := financial.Currency("EUR")

		want := financial.Money{}

		got, err := financial.NewMoney(amount, currency)

		if err == nil || got != want {
			t.Errorf("money.NewMoney(%d, %s) = %v %v; want %v error\n%v", amount, currency, got, err, want, cmp.Diff(want, got))
		}
	})
}
