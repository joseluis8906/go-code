package delivery

// Address represents an address.
type Address struct {
	Street string
	City   string
}

// NewAddress creates a new address.
func NewAddress(street, city string) Address {
	return Address{
		Street: street,
		City:   city,
	}
}
