package domain

import (
	"context"
	"time"
)

type IntervalScrapper interface {
	Results() <-chan ScrapResult
	SetInterval(interval time.Duration)
	Stop()
}

type intervalScrapper struct {
	ctx         context.Context
	cancelCtx   context.CancelFunc
	scrapper    Scrapper
	ticker      *time.Ticker
	chanResults chan ScrapResult
}

func NewIntervalScrapper(scrapper Scrapper, interval time.Duration) IntervalScrapper {
	isc := intervalScrapper{
		ctx:         nil,
		cancelCtx:   nil,
		scrapper:    scrapper,
		ticker:      time.NewTicker(interval),
		chanResults: make(chan ScrapResult),
	}

	isc.start()

	go func() {
		isc.executeScrapper()
	}()

	return &isc
}

func (isc *intervalScrapper) start() {
	isc.ctx, isc.cancelCtx = context.WithCancel(context.Background())

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
}

func (isc *intervalScrapper) executeScrapper() {
	isc.chanResults <- <-isc.scrapper.Scrap(isc.ctx)
}

func (isc *intervalScrapper) Results() <-chan ScrapResult {
	return isc.chanResults
}

func (isc *intervalScrapper) SetInterval(interval time.Duration) {
	isc.Stop()
	isc.ticker.Reset(interval)
	isc.start()
}

func (isc *intervalScrapper) Stop() {
	isc.ticker.Stop()
	isc.cancelCtx()
}
