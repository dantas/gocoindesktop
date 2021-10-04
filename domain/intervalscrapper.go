package domain

import (
	"context"
	"time"
)

type IntervalScrapper interface {
	Coins() <-chan []Coin
	SetInterval(interval time.Duration)
	Destroy()
}

type implIntervalScrapper struct {
	cancelCtx context.CancelFunc
	ticker    *time.Ticker
	channel   chan []Coin
}

func NewIntervalScrapper(scrapper Scrapper, interval time.Duration) IntervalScrapper {
	ctx, cancelCtx := context.WithCancel(context.Background())

	perScrapper := implIntervalScrapper{
		cancelCtx: cancelCtx,
		ticker:    time.NewTicker(interval),
		channel:   make(chan []Coin),
	}

	go func() {
		perScrapper.channel <- Scrap(ctx, scrapper)
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(perScrapper.channel)
				return
			case <-perScrapper.ticker.C:
				perScrapper.channel <- Scrap(ctx, scrapper)
			}
		}
	}()

	return &perScrapper
}

func (isc *implIntervalScrapper) Coins() <-chan []Coin {
	return isc.channel
}

func (isc *implIntervalScrapper) SetInterval(interval time.Duration) {
	if isc.ticker != nil {
		isc.ticker.Reset(interval)
	}
}

func (isc *implIntervalScrapper) Destroy() {
	isc.cancelCtx()
	if isc.ticker != nil {
		isc.ticker.Stop()
		isc.ticker = nil
	}
}
