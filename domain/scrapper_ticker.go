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

				for c := range scrapper(tmScrapper.done) {
					coins = append(coins, c)
				}

				tmScrapper.channel <- coins
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
