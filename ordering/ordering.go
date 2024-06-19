package ordering

const (
	// Less is a constant to represent a comparison result of -1.
	Less = Cmp(-1)
	// Equal is a constant to represent a comparison result of 0.
	Equal = Cmp(0)
	// Greater is a constant to represent a comparison result of 1.
	Greater = Cmp(1)
)

type (
	// Cmp is a type alias for int.
	Cmp int
)
