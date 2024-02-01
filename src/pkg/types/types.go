package types

type (
	StringValue struct {
		Value string
		Valid bool
	}

	IntValue struct {
		Value int
		Valid bool
	}

	UintValue struct {
		Value uint
		Valid bool
	}

	FloatValue struct {
		Value float32
		Valid bool
	}

	Float64Value struct {
		Value float64
		Valid bool
	}
)
