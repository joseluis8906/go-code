package financial

import (
	"errors"
	"math"
)

type (
	// Bank represents a bank that can convert currencies using rates between them.
	// It is not thread safe.
	Bank struct {
		rates map[Pair]float64
	}

	// Pair represents a pair of currencies.
	Pair struct {
		from Currency
		to   Currency
	}
)

// NewBank creates a new Bank instance.
func NewBank() Bank {
	return Bank{
		rates: map[Pair]float64{},
	}
}

// Reduce reduces a Money to a given currency.
// If the currency is the same as the Money currency, it returns the same Money.
// If the currency is different, it uses the rates between the currencies to calculate the new Money.
//
//	bank := NewBank()
//	bank.AddRate("CHF", "USD", 2)
//	result := bank.Reduce(Franc(2), "USD") // result is Dollar(1)
func (b Bank) Reduce(from Money, to Currency) (Money, error) {
	if from.Currency == to {
		return from, nil
	}

	pair := Pair{from.Currency, to}
	rate, ok := b.rates[pair]

	if !ok {
		return from, errors.New("rate not found")
	}

	return NewMoney(int64(math.Round(float64(from.Amount)*rate)), to)
}

// AddRate adds a rate between two currencies.
func (b *Bank) AddRate(from, to Currency, rate float64) {
	pair := Pair{from, to}
	b.rates[pair] = rate

	inversedPair := Pair{to, from}
	b.rates[inversedPair] = 1 / rate
}

// Rate returns the rate between two currencies. If the rate is not found, it panics.
// Use RateOk to check if a rate exists.
func (b Bank) Rate(from, to Currency) float64 {
	pair := Pair{from, to}
	rate, ok := b.rates[pair]
	if !ok {
		return 1
	}

	return rate
}
