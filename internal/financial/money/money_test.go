package money_test

import (
	"testing"

	"github.com/joseluis8906/go-code/internal/financial/bank"
	"github.com/joseluis8906/go-code/internal/financial/money"
	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := money.Dollar(5)
	assert.Equal(t, money.Dollar(10), five.Times(2))
	assert.Equal(t, money.Dollar(15), five.Times(3))
}

func TestEquality(t *testing.T) {
	assert.True(t, money.Dollar(5).Equals(money.Dollar(5)))
	assert.False(t, money.Dollar(5).Equals(money.Dollar(6)))
	assert.True(t, money.Franc(5).Equals(money.Franc(5)))
}

func TestPlus(t *testing.T) {
	five := money.Dollar(5)
	assert.Equal(t, money.Dollar(10), five.Plus(money.Dollar(5)))
	assert.Zero(t, five.Plus(money.Franc(5)))
}

func TestReduceDifferentCurrency(t *testing.T) {
	itau := bank.New()
	itau.AddRate(money.CHF, money.USD, 0.5)
	assert.Equal(t, money.Dollar(1), itau.Reduce(money.Franc(2), money.USD))
	assert.Equal(t, money.Franc(10), itau.Reduce(money.Dollar(5), money.CHF))
	assert.Equal(t, money.Dollar(6), itau.Reduce(money.Dollar(6), money.USD))
}

func TestValidCurrencies(t *testing.T) {
	assert.Zero(t, money.New(100, "UNKNOWN"))
	assert.Equal(t, money.Dollar(3), money.New(3, money.USD))
}
