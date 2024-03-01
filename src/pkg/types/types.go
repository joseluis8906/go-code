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

func (sv StringValue) IsZero() bool {
	return sv.Value == ""
}

func (iv IntValue) IsZero() bool {
	return iv.Value == 0
}

func (uv UintValue) IsZero() bool {
	return uv.Value == 0
}

func (fv FloatValue) IsZero() bool {
	return fv.Value == 0
}

func (fv Float64Value) IsZero() bool {
	return fv.Value == 0
}
