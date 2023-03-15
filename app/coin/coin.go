package coin

type Coin struct {
	Name  string
	Price float64
}

type CoinSource func() ([]Coin, error)
