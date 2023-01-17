package bank

import (
	"math"

	"github.com/joseluis8906/go-code/internal/financial"
	"github.com/joseluis8906/go-code/internal/financial/money"
)

type (
	Bank struct {
		rates map[Pair]float64
	}

	Pair struct {
		from string
		to   string
	}
)

func New() Bank {
	return Bank{
		rates: map[Pair]float64{},
	}
}

func (b Bank) Reduce(source financial.ExchangeUnit, to string) financial.ExchangeUnit {
  if source.Currency() == to {
    return source
  }

	pair := Pair{source.Currency(), to}
	rate, ok := b.rates[pair]
	if !ok {
		panic("rate not found")
	}

	return money.New(int(math.Round(float64(source.Amount())*rate)), to)
}

func (b *Bank) AddRate(from, to string, rate float64) {
	pair := Pair{from, to}
	b.rates[pair] = rate

  inversedPair := Pair{to, from}
  b.rates[inversedPair] = 1/rate
}
