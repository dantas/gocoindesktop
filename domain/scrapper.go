package domain

type Scrapper func(done <-chan interface{}) <-chan ScrapResult

type ScrapResult struct {
	Coin  Coin
	Error error
}
