package financial_test

import (
	"testing"

	"github.com/joseluis8906/go-code/src/pkg/financial"
	"github.com/stretchr/testify/assert"
)

func TestBank_Reduce(t *testing.T) {
	t.Parallel()

	centralBank := financial.NewBank()
	centralBank.AddRate(financial.CHF, financial.USD, 0.5)

	assert.Equal(t, financial.Dollar(1), centralBank.Reduce(financial.Franc(2), financial.USD))
	assert.Equal(t, financial.Franc(10), centralBank.Reduce(financial.Dollar(5), financial.CHF))
	assert.Equal(t, financial.Dollar(6), centralBank.Reduce(financial.Dollar(6), financial.USD))
}

func TestBank_Rate(t *testing.T) {
	t.Parallel()

	centralBank := financial.NewBank()
	centralBank.AddRate(financial.CHF, financial.USD, 0.5)

	assert.InEpsilon(t, 0.5, centralBank.Rate(financial.CHF, financial.USD), 0.1)
	assert.InEpsilon(t, 2, centralBank.Rate(financial.USD, financial.CHF), 0.1)
	assert.InEpsilon(t, 1, centralBank.Rate(financial.USD, financial.USD), 0.1)
}
