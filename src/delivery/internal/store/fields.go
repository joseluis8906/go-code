package store

type Name struct {
	value string
}

func (n Name) String() string {
	return n.value
}
