package domain

import "errors"

type Coin struct {
	Name  string
	Price float64
}

type CoinSource func() ([]Coin, error)

var ErrCoinSource = errors.New("error fetching coins from coin source")
