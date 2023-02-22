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

type intervalScrapper struct {
	ctx         context.Context
	cancelCtx   context.CancelFunc
	scrapper    Scrapper
	ticker      *time.Ticker
	chanResults chan ScrapResult
}

func NewIntervalScrapper(scrapper Scrapper) IntervalScrapper {
	ctx, cancelCtx := context.WithCancel(context.Background())

	isc := intervalScrapper{
		ctx:         ctx,
		cancelCtx:   cancelCtx,
		scrapper:    scrapper,
		ticker:      time.NewTicker(time.Second),
		chanResults: make(chan ScrapResult),
	}

	isc.ticker.Stop()

	go func() {
		for {
			select {
			case <-isc.ctx.Done():
				close(isc.chanResults)
				return
			case <-isc.ticker.C:
				isc.executeScrapper()
			}
		}
	}()

	go func() {
		isc.executeScrapper()
	}()

	return &isc
}

func (isc *intervalScrapper) executeScrapper() {
	isc.chanResults <- <-isc.scrapper.Scrap(isc.ctx)
}

func (isc *intervalScrapper) Results() <-chan ScrapResult {
	return isc.chanResults
}

func (isc *intervalScrapper) SetInterval(interval time.Duration) {
	isc.ticker.Reset(interval)
}

func (isc *intervalScrapper) Destroy() {
	isc.ticker.Stop()
	isc.cancelCtx()
}
