package financial

import "errors"

// Money is a value object that represents a certain amount of money in a certain currency.
// It is immutable. It can be added, subtracted, multiplied and compared.
// It can be created using the Dollar() and Franc() functions.

const (
	// USD is the currency code for US Dollars.
	USD Currency = "USD"
	// CHF is the currency code for Swiss Francs.
	CHF Currency = "CHF"
)

type (
	// Money is a value object that represents a certain amount of money in a certain currency.
	// It is immutable. It can be added, subtracted, multiplied and compared.
	// It can be created using the Dollar() and Franc() functions.
	//
	//  five := Dollar(5)
	//  ten := five.Times(2)
	Money struct {
		Amount   int64
		Currency Currency
	}

	Currency string
)

// NewMoney creates a new Money object.
// It returns a Noop() if the currency is not supported.
//
//	five := New(5, "USD")
func NewMoney(amount int64, currency Currency) (Money, error) {
	switch currency {
	case USD, CHF:
	default:
		return Noop(), errors.New("unsupported currency")
	}

	return Money{amount, currency}, nil
}

// Noop returns a Money object with a zero amount and an empty currency.
// It is used to indicate that a Money object could not be created.
//
//	five := New(5, "EUR") // five is a Noop()
func Noop() Money {
	return Money{0, ""}
}

// Dollar creates a new Money object with the currency set to USD. It is a shortcut for New(5, "USD").
// It can be used like this:
//
//	five := Dollar(5)
func Dollar(amount int64) Money {
	return Money{amount, "USD"}
}

// Franc creates a new Money object with the currency set to CHF. It is a shortcut for New(5, "CHF").
// It can be used like this:
//
//	five := Franc(5)
func Franc(amount int64) Money {
	return Money{amount, "CHF"}
}

// Equals compares two Money objects and returns true if they have the same amount and currency.
// It can be used like this:
//
//	five := Dollar(5)
//	five.Equals(ten) // false
func (m Money) Equals(another Money) bool {
	return m.Amount == another.Amount && m.Currency == another.Currency
}

// Times multiplies the amount of the Money object by the given multiplier and returns a new Money object.
// It can be used like this:
//
//	five := Dollar(5)
//	ten := five.Times(2) // ten is a Dollar(10)
func (m Money) Times(multiplier int64) Money {
	return Money{m.Amount * multiplier, m.Currency}
}

// Plus adds the amount of the given Money object to the amount of the Money object and returns a new Money object.
// It can be used like this:
//
//	five := Dollar(5)
//	ten := five.Plus(five) // ten is a Dollar(10)
func (m Money) Plus(addend Money) Money {
	if m.Currency != addend.Currency {
		return Money{0, ""}
	}

	return Money{m.Amount + addend.Amount, m.Currency}
}
