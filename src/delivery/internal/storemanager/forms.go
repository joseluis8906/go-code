package storemanager

type (
	StoreForm struct {
		Name     string
		Country  string
		City     string
		Address  string
		Products []ProductForm
	}

	ProductForm struct {
		Ref   string
		Name  string
		Price MoneyForm
	}

	MoneyForm struct {
		Amount   int64
		Currency string
	}
)
