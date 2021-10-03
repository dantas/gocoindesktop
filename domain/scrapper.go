package domain

import "context"

type Scrapper func(ctx context.Context) <-chan ScrapResult

type ScrapResult struct {
	Coin  Coin
	Error error
}

func CollectScrapperResults(resultsChannel <-chan ScrapResult) []Coin {
	coins := make([]Coin, 0)

	for result := range resultsChannel {
		// We operate on best effort, we collect any coin available
		if result.Error == nil {
			coins = append(coins, result.Coin)
		}
	}

	return coins
}
