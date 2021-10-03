package domain

import (
	"context"
	"time"
)

type ScrapperTicker interface {
	Channel() <-chan []Coin
	SetInterval(interval time.Duration)
	Stop()
}

type scrapperTicker struct {
	cancelCtx context.CancelFunc
	ticker    *time.Ticker
	channel   chan []Coin
}

func NewScrapperTicker(scrapper Scrapper, interval time.Duration) ScrapperTicker {
	ctx, cancelCtx := context.WithCancel(context.Background())

	ticker := scrapperTicker{
		cancelCtx: cancelCtx,
		ticker:    time.NewTicker(interval),
		channel:   make(chan []Coin),
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(ticker.channel)
				return
			case <-ticker.ticker.C:
				resultsChannel := scrapper(ctx)
				coins := CollectScrapperResults(resultsChannel)
				ticker.channel <- coins
			}
		}
	}()

	return &ticker
}

func (st *scrapperTicker) Channel() <-chan []Coin {
	return st.channel
}

func (st *scrapperTicker) SetInterval(interval time.Duration) {
	st.ticker.Reset(interval)
}

func (st *scrapperTicker) Stop() {
	st.cancelCtx()
	st.ticker.Stop()
}
