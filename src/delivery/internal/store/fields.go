package store

import (
	"fmt"

	"github.com/joseluis8906/go-code/src/pkg/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type fields struct {
	Name     string
	Products struct {
		Ref  string
		Name string
	}
}

func Fields() fields {
	var f fields
	f.Name = "name"
	f.Products.Ref = "products.ref"
	f.Products.Name = "products.name"

	return f
}

type (
	Name    types.StringValue
	Country types.StringValue
	City    types.StringValue
	Address types.StringValue
)

func NewName(value string) (Name, error) {
	if len(value) < 2 {
		return Name{}, fmt.Errorf("name length must be bigger than 2")
	}

	return Name{Value: value, Valid: true}, nil
}

func NewCountry(value string) (Country, error) {
	return Country{Value: value, Valid: true}, nil
}

func NewCity(value string) (City, error) {
	return City{Value: value, Valid: true}, nil
}

func NewAddress(value string) (Address, error) {
	return Address{Value: value, Valid: true}, nil
}

func (n Name) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(n.Value)
}

func (n *Name) UnmarshalBSONValue(t bsontype.Type, val []byte) error {
	if t != bson.TypeString {
		return fmt.Errorf("invalid bson type '%s'", t.String())
	}

	if err := bson.UnmarshalValue(t, val, &n.Value); err != nil {
		return fmt.Errorf("unmarshaling value: %w", err)
	}

	n.Valid = true

	return nil
}

func (n Name) String() string {
	return n.Value
}

func (c Country) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(c.Value)
}

func (c *Country) UnmarshalBSONValue(t bsontype.Type, val []byte) error {
	if t != bson.TypeString {
		return fmt.Errorf("invalid bson type '%s'", t.String())
	}

	if err := bson.UnmarshalValue(t, val, &c.Value); err != nil {
		return fmt.Errorf("unmarshaling value: %w", err)
	}

	c.Valid = true

	return nil
}

func (c City) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(c.Value)
}

func (c *City) UnmarshalBSONValue(t bsontype.Type, val []byte) error {
	if t != bson.TypeString {
		return fmt.Errorf("invalid bson type '%s'", t.String())
	}

	if err := bson.UnmarshalValue(t, val, &c.Value); err != nil {
		return fmt.Errorf("unmarshaling value: %w", err)
	}

	c.Valid = true

	return nil
}

func (c Address) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(c.Value)
}

func (c *Address) UnmarshalBSONValue(t bsontype.Type, val []byte) error {
	if t != bson.TypeString {
		return fmt.Errorf("invalid bson type '%s'", t.String())
	}

	if err := bson.UnmarshalValue(t, val, &c.Value); err != nil {
		return fmt.Errorf("unmarshaling value: %w", err)
	}

	c.Valid = true

	return nil
}

func (n Name) IsZero() bool {
	return n.Value == ""
}
