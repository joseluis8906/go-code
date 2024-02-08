package product

import "context"

type (
	// Product is an extended delivery product.
	Product struct {
		Ref   Ref
		Name  Name
		Price Price
	}

	productBuilder struct {
		product Product
		err     error
	}
)

func New() *productBuilder {
	return &productBuilder{}
}

func (pb *productBuilder) Ref(ref string) *productBuilder {
	if pb.err != nil {
		return pb
	}

	pb.product.Ref, pb.err = NewRef(ref)
	return pb
}

func (pb *productBuilder) Name(name string) *productBuilder {
	if pb.err != nil {
		return pb
	}

	pb.product.Name, pb.err = NewName(name)
	return pb
}

func (pb *productBuilder) Price(amount int64, currency string) *productBuilder {
	if pb.err != nil {
		return pb
	}

	pb.product.Price, pb.err = NewPrice(amount, currency)
	return pb
}

func (pb *productBuilder) Do(ctx context.Context) (Product, error) {
	if pb.err != nil {
		return Product{}, pb.err
	}

	return pb.product, nil
}
