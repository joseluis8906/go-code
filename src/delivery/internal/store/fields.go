package store

type (
	Name struct {
		value string
	}

	City struct {
		value string
	}

	Address struct {
		value string
	}
)

func (n Name) String() string {
	return n.value
}

func (c City) String() string {
	return c.value
}

func (a Address) String() string {
	return a.value
}
