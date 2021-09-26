package domain

import (
	"time"
)

type ScrapperTicker interface {
	Channel() <-chan []Coin
	SetInterval(interval time.Duration)
	Stop()
}

type scrapperTicker struct {
	done    chan interface{}
	ticker  *time.Ticker
	channel chan []Coin
}

func NewScrapperTicker(scrapper Scrapper, interval time.Duration) ScrapperTicker {
	tmScrapper := scrapperTicker{
		done:    make(chan interface{}),
		ticker:  time.NewTicker(interval),
		channel: make(chan []Coin),
	}

	go func() {
		for {
			select {
			case <-tmScrapper.done:
				close(tmScrapper.channel)
				return
			case <-tmScrapper.ticker.C:
				coins := make([]Coin, 0)

				for result := range scrapper(tmScrapper.done) {
					// We operate on best effort, we attempt to collect any coin available
					if result.Error == nil {
						coins = append(coins, result.Coin)
					}
				}

				// Only bother caller if we got any coin
				if len(coins) > 0 {
					tmScrapper.channel <- coins
				}
			}

		}
	}()

	return &tmScrapper
}

func (st *scrapperTicker) Channel() <-chan []Coin {
	return st.channel
}

func (st *scrapperTicker) SetInterval(interval time.Duration) {
	st.ticker.Reset(interval)
}

func (st *scrapperTicker) Stop() {
	close(st.done)
	st.ticker.Stop()
}
