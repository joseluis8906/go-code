package delivery

// Address represents an address.
type Address struct {
	Name    string
	City    string
	Country string
}

// NewAddress creates a new address.
func NewAddress(name, city, country string) Address {
	return Address{
		Name:    name,
		City:    city,
		Country: country,
	}
}
