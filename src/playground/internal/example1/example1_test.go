package example1_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 1, 1)
}
