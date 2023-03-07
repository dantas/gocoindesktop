package domain

import (
	"time"
)

type coinTicker struct {
	done   chan struct{}
	source CoinSource
	ticker *time.Ticker
	Coins  chan []Coin
	Errors chan error
}

func NewCoinTicker(source CoinSource) *coinTicker {
	ct := coinTicker{
		done:   make(chan struct{}),
		source: source,
		ticker: time.NewTicker(time.Second),
		Coins:  make(chan []Coin),
		Errors: make(chan error),
	}

	ct.ticker.Stop()

	go func() {
		for {
			select {
			case <-ct.done:
				close(ct.Coins)
				close(ct.Errors)
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
		ct.Errors <- err
	} else {
		ct.Coins <- coins
	}
}

func (ct *coinTicker) SetInterval(interval time.Duration) {
	ct.ticker.Reset(interval)
}

func (ct *coinTicker) Destroy() {
	close(ct.done)
}
