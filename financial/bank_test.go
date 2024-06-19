package financial_test

import (
	"testing"

	"github.com/joseluis8906/yummies/go-code/src/pkg/cmp"
	"github.com/joseluis8906/yummies/go-code/src/pkg/financial"
)

func TestBank_Reduce(t *testing.T) {
	t.Parallel()

	centralBank := financial.NewBank()
	centralBank.AddRate(financial.CHF, financial.USD, 0.5)

	testCases := map[string]struct {
		from financial.Money
		to   financial.Currency
		want financial.Money
	}{
		"FrancToDollar": {
			from: financial.Franc(2),
			to:   financial.USD,
			want: financial.Dollar(1),
		},
		"DollarToFranc": {
			from: financial.Dollar(5),
			to:   financial.CHF,
			want: financial.Franc(10),
		},
		"DollarToDollar": {
			from: financial.Dollar(6),
			to:   financial.USD,
			want: financial.Dollar(6),
		},
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := centralBank.Reduce(tc.from, tc.to)

			if err != nil || got != tc.want {
				t.Errorf("bank.Reduce(%v, %s) = %v, %v; want %v, nil\n%v", tc.from, tc.to, got, err, tc.want, cmp.Diff(tc.want, got))
			}
		})
	}
}

func TestBank_Rate(t *testing.T) {
	t.Parallel()

	centralBank := financial.NewBank()
	centralBank.AddRate(financial.CHF, financial.USD, 0.5)

	testCases := map[string]struct {
		from financial.Currency
		to   financial.Currency
		want float64
	}{
		"CHFToUSD": {
			from: financial.CHF,
			to:   financial.USD,
			want: 0.5,
		},
		"USDToCHF": {
			from: financial.USD,
			to:   financial.CHF,
			want: 2,
		},
		"USDToUSD": {
			from: financial.USD,
			to:   financial.USD,
			want: 1,
		},
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := centralBank.Rate(tc.from, tc.to)

			if got != tc.want {
				t.Errorf("bank.Rate(%s, %s) = %f; want %f\n%v", tc.from, tc.to, got, tc.want, cmp.Diff(tc.want, got))
			}
		})
	}
}
