package financial_test

import (
	"testing"

	"github.com/joseluis8906/go-code/src/financial"
	"github.com/stretchr/testify/assert"
)

func TestMoney_Multiplication(t *testing.T) {
	t.Parallel()

	five := financial.Dollar(5)

	assert.Equal(t, financial.Dollar(10), five.Times(2))
	assert.Equal(t, financial.Dollar(15), five.Times(3))
}

func TestMoney_Equality(t *testing.T) {
	t.Parallel()

	assert.True(t, financial.Dollar(5).Equals(financial.Dollar(5)))
	assert.False(t, financial.Dollar(5).Equals(financial.Dollar(6)))
	assert.True(t, financial.Franc(5).Equals(financial.Franc(5)))
}

func TestMoney_Plus(t *testing.T) {
	t.Parallel()

	five := financial.Dollar(5)

	assert.Equal(t, financial.Dollar(10), five.Plus(financial.Dollar(5)))
	assert.Zero(t, five.Plus(financial.Franc(5)))
}

func TestMoney_ValidCurrencies(t *testing.T) {
	t.Parallel()

	assert.Zero(t, financial.NewMoney(100, "UNKNOWN"))
	assert.Equal(t, financial.Dollar(3), financial.NewMoney(3, financial.USD))
}
