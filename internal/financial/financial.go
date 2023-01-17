package financial

type (
	ExchangeUnit interface {
		Amount() int
		Currency() string
	}
)
