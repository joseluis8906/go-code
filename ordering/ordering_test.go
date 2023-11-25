package ordering_test

import (
	"testing"

	"github.com/joseluis8906/go-code/ordering"
	"github.com/stretchr/testify/assert"
)

func TestOrdering(t *testing.T) {
	t.Parallel()

	assert.Less(t, ordering.Less, ordering.Equal)
	assert.Less(t, ordering.Equal, ordering.Greater)
}
