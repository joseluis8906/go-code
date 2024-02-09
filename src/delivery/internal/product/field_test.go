package product

import (
	"testing"

	"github.com/joseluis8906/go-code/src/pkg/cmp"
	"go.mongodb.org/mongo-driver/bson"
)

func TestName(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		t.Parallel()

		in := "Hmaburger"
		want := Name{in, true}

		got, err := NewName(in)

		if got != want || err != nil {
			t.Errorf("NewName(%v) = %v, %v; want %v, nil\n%v", in, got, err, want, cmp.Diff(want, got))
		}
	})

	t.Run("TooShort", func(t *testing.T) {
		t.Parallel()

		in := "Ha"
		var want Name

		got, err := NewName(in)

		if got != want || err == nil {
			t.Errorf("NewName(%v) = %v, %v; want %v, error\n%v", in, got, err, want, cmp.Diff(want, got))
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
			t.Errorf("NewRef(%v) = %v, %v; want %v, nil\n%v", in, got, err, want, cmp.Diff(want, got))
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
					t.Errorf("NewRef(%v) = %v, %v; want %v, error\n%v", tc.in, got, err, tc.want, cmp.Diff(tc.want, got))
				}
			})
		}
	})
}

func TestPrice(t *testing.T) {
	t.Run("MarshalBSON", func(t *testing.T) {
		t.Parallel()

		in := struct {
			Amount   int64
			Currency string
		}{
			Amount:   1000,
			Currency: "USD",
		}

		_, want, _ := bson.MarshalValue(in)
		price, _ := NewPrice(in.Amount, in.Currency)

		got, err := price.MarshalBSON()

		if err != nil || !cmp.Equal(got, want) {
			t.Errorf("Price.MarshalBSON() = %v, %v; want %v, nil\n%v", got, err, want, cmp.Diff(want, got))
		}
	})

	t.Run("UnmarshalBSON", func(t *testing.T) {
		t.Parallel()

		in := struct {
			Amount   int64
			Currency string
		}{
			Amount:   1000,
			Currency: "USD",
		}

		_, data, _ := bson.MarshalValue(in)

		want, _ := NewPrice(in.Amount, in.Currency)

		var got Price
		err := got.UnmarshalBSON(data)

		if err != nil || !cmp.Equal(got, want) {
			t.Errorf("Price.UnmarshalBSON(%v) = %v, %v; want %v, <nil>\n%v", data, got, err, want, cmp.Diff(want, got))
		}
	})
}
