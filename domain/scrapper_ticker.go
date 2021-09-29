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
	ticker := scrapperTicker{
		done:    make(chan interface{}),
		ticker:  time.NewTicker(interval),
		channel: make(chan []Coin),
	}

	go func() {
		for {
			select {
			case <-ticker.done:
				close(ticker.channel)
				return
			case <-ticker.ticker.C:
				resultsChannel := scrapper(ticker.done)
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
	close(st.done)
	st.ticker.Stop()
}