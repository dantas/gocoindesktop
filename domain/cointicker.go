package domain

import (
	"time"
)

type CoinTicker struct {
	done   chan struct{}
	source CoinSource
	ticker *time.Ticker
	coins  chan []Coin
	errors chan error
}

func NewCoinTicker(source CoinSource) *CoinTicker {
	ct := CoinTicker{
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

func (ct *CoinTicker) fetchCoins() {
	coins, err := ct.source()

	if err != nil {
		ct.errors <- err
	} else {
		ct.coins <- coins
	}
}

func (ct *CoinTicker) Coins() <-chan []Coin {
	return ct.coins
}

func (ct *CoinTicker) Errors() <-chan error {
	return ct.errors
}

func (ct *CoinTicker) SetInterval(interval time.Duration) {
	ct.ticker.Reset(interval)
}

func (ct *CoinTicker) Destroy() {
	close(ct.done)
}
