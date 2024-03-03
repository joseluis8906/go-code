package product

import "github.com/joseluis8906/go-code/protobuf/deliverypb"

type (
	// Product is an extended delivery product.
	Product struct {
		Ref   Ref
		Name  Name
		Price Price
	}

	Builder struct {
		product Product
		err     error
	}
)

func New() *Builder {
	return &Builder{}
}

func (pb *Builder) Ref(ref string) *Builder {
	if pb.err != nil {
		return pb
	}

	pb.product.Ref, pb.err = NewRef(ref)
	return pb
}

func (pb *Builder) Name(name string) *Builder {
	if pb.err != nil {
		return pb
	}

	pb.product.Name, pb.err = NewName(name)
	return pb
}

func (pb *Builder) Price(amount int64, currency string) *Builder {
	if pb.err != nil {
		return pb
	}

	pb.product.Price, pb.err = NewPrice(amount, currency)
	return pb
}

func (pb *Builder) Build() (Product, error) {
	if pb.err != nil {
		return Product{}, pb.err
	}

	return pb.product, nil
}

func FromPB(data *deliverypb.Product) (Product, error) {
	ref := data.GetRef().GetValue()
	name := data.GetName().GetValue()
	amount := data.GetPrice().GetAmount().GetValue()
	currency := data.GetPrice().GetCurrency().GetValue()

	var pb Builder
	pb.Ref(ref)
	pb.Name(name)
	pb.Price(amount, currency)

	return pb.Build()
}
