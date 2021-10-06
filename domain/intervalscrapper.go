package domain

import (
	"context"
	"time"
)

type IntervalScrapper interface {
	Results() <-chan ScrapResult
	SetInterval(interval time.Duration)
	Destroy()
}

type implIntervalScrapper struct {
	cancelCtx context.CancelFunc
	ticker    *time.Ticker
	results   chan ScrapResult
}

func NewIntervalScrapper(scrapper Scrapper, interval time.Duration) IntervalScrapper {
	ctx, cancelCtx := context.WithCancel(context.Background())

	intScrapper := implIntervalScrapper{
		cancelCtx: cancelCtx,
		ticker:    time.NewTicker(interval),
		results:   make(chan ScrapResult),
	}

	go func() {
		intScrapper.results <- <-scrapper(ctx)
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(intScrapper.results)
				return
			case <-intScrapper.ticker.C:
				intScrapper.results <- <-scrapper(ctx)
			}
		}
	}()

	return &intScrapper
}

func (is *implIntervalScrapper) Results() <-chan ScrapResult {
	return is.results
}

func (is *implIntervalScrapper) SetInterval(interval time.Duration) {
	if is.ticker != nil {
		is.ticker.Reset(interval)
	}
}

func (is *implIntervalScrapper) Destroy() {
	is.cancelCtx()
	if is.ticker != nil {
		is.ticker.Stop()
		is.ticker = nil
	}
}
