package domain_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/dantas/gocoindesktop/domain"
)

var results = []domain.ScrapResult{
	{
		Coin: domain.Coin{
			Name: "first coin",
		},
	},
	{
		Error: errors.New("something happened"),
	},
	{
		Coin: domain.Coin{
			Name: "second coin",
		},
	},
	{
		Coin: domain.Coin{
			Name: "third coin",
		},
	},
}

func TestScrapperTicker(t *testing.T) {
	scrapperTicker := domain.NewIntervalScrapper(createMockScrapper(), 1*time.Second)

	index := 0

	for coins := range scrapperTicker.Coins() {
		if index == 1 {
			if len(coins) != 0 {
				t.Error("Ticker returned a coin where an error was expected, index", index)
			}
		} else {
			if len(coins) != 1 {
				t.Errorf("Scrapper returning different amount of values than what is expected %#v\n", coins)
			}

			if coins[0] != results[index].Coin {
				t.Error("Scrapper didn`t return expected result, index", index)
			}
		}

		index += 1

		if index == 3 {
			scrapperTicker.Destroy()
		}
	}
}

func createMockScrapper() domain.Scrapper {
	index := 0

	mockScrapper := func(context.Context) <-chan domain.ScrapResult {
		resultChannel := make(chan domain.ScrapResult)

		go func() {
			resultChannel <- results[index]

			index += 1

			close(resultChannel)
		}()

		return resultChannel
	}

	return mockScrapper
}
