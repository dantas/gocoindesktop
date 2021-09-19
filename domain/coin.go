package domain

type Coin struct {
	Name  string
	Price float64
}

type Scrapper func(done <-chan interface{}) <-chan Coin
