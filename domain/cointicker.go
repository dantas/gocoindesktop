package domain

import (
	"time"
)

type CoinTicker interface {
	Coins() <-chan []Coin
	Errors() <-chan error
	SetInterval(interval time.Duration)
	Destroy()
}

type coinTicker struct {
	done   chan struct{}
	source CoinSource
	ticker *time.Ticker
	coins  chan []Coin
	errors chan error
}

func NewCoinTicker(source CoinSource) CoinTicker {
	ct := coinTicker{
		done:   make(chan struct{}),
		source: source,
		ticker: time.NewTicker(time.Second),
		coins:  make(chan []Coin),
		errors: make(chan error),
	}

	ct.ticker.Stop()

	go func() {
		for {
			select {
			case <-ct.done:
				close(ct.coins)
				close(ct.errors)
				return
			case <-ct.ticker.C:
				ct.fetchCoins()
			}
		}
	}()

	go func() {
		ct.fetchCoins()
	}()

	return &ct
}

func (ct *coinTicker) fetchCoins() {
	coins, err := ct.source()

	if err != nil {
		ct.errors <- err
	} else {
		ct.coins <- coins
	}
}

func (ct *coinTicker) Coins() <-chan []Coin {
	return ct.coins
}

func (ct *coinTicker) Errors() <-chan error {
	return ct.errors
}

func (ct *coinTicker) SetInterval(interval time.Duration) {
	ct.ticker.Reset(interval)
}

func (ct *coinTicker) Destroy() {
	close(ct.done)
}
