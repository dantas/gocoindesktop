package domain

import (
	"context"
	"fmt"
)

// =======
// Data layer implements this function

type CoinSource func(ctx context.Context) <-chan CoinSourceResult

type CoinSourceResult struct {
	Coin  Coin
	Error error
}

// =======

type Coin struct {
	Name  string
	Price float64
}

type ScrapResult struct {
	Coins []Coin
	Error error
}

type Scrapper interface {
	Scrap(ctx context.Context) <-chan ScrapResult
}

func NewScrapper(coinSource CoinSource) scrapper {
	return scrapper{
		coinSource: coinSource,
	}
}

type scrapper struct {
	coinSource CoinSource
}

func (s scrapper) Scrap(ctx context.Context) <-chan ScrapResult {
	chanResult := make(chan ScrapResult)

	go func() {
		coins := make([]Coin, 0)
		var err error

		for r := range s.coinSource(ctx) {
			if r.Error != nil {
				err = fmt.Errorf("%w", r.Error)
			} else {
				coins = append(coins, r.Coin)
			}
		}

		chanResult <- ScrapResult{
			Coins: coins,
			Error: err,
		}
	}()

	return chanResult
}
