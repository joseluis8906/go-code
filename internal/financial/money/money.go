package money

const (
	USD = "USD"
	CHF = "CHF"
)

type (
	Money struct {
		amount   int
		currency string
	}
)

func new(amount int, currency string) Money {
	return Money{
		amount:   amount,
		currency: currency,
	}
}

func New(amount int, currency string) Money {
	switch currency {
	case USD, CHF:
	default:
		return Noop()
	}

	return new(amount, currency)
}

func Noop() Money {
  return new(0, "")
}

func Dollar(amount int) Money {
	return new(amount, "USD")
}

func Franc(amount int) Money {
	return new(amount, "CHF")
}

func (m Money) Equals(another Money) bool {
	return m.Amount() == another.Amount() && m.Currency() == another.Currency()
}

func (m Money) Currency() string {
	return m.currency
}

func (m Money) Amount() int {
	return m.amount
}

func (m Money) Times(multiplier int) Money {
	return new(m.amount*multiplier, m.currency)
}

func (m Money) Plus(addend Money) Money {
	if m.Currency() != addend.Currency() {
		return new(0, "")
	}

	return new(m.Amount()+addend.Amount(), m.Currency())
}
