package domain_test

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/dantas/gocoindesktop/domain"
)

var scrapResults = []domain.ScrapResult{
	{
		Coins: []domain.Coin{
			{
				Name:  "First Coin",
				Price: 12,
			},
			{
				Name:  "Answer Coin",
				Price: 42,
			},
			{
				Name:  "Nice Coin",
				Price: 69,
			},
		},
		Error: nil,
	},
	{
		Coins: []domain.Coin{},
		Error: errors.New("tragic error"),
	},
}

func TestIntervalScrapper(t *testing.T) {
	scrapper := domain.NewScrapper(createCoinMockSource())
	intervalScrapper := domain.NewIntervalScrapper(scrapper)

	intervalScrapper.SetInterval(1 * time.Second)

	index := 0

	for result := range intervalScrapper.Results() {
		if !reflect.DeepEqual(result.Coins, scrapResults[index].Coins) {
			t.Errorf("Returned result is different from what is expected %v != %v\n", result, scrapResults[index])
		}

		if !errors.Is(result.Error, scrapResults[index].Error) {
			t.Errorf("Returned error is different from what is expected %v != %v\n", result.Error, scrapResults[index].Error)
		}

		index += 1

		if index == len(scrapResults) {
			intervalScrapper.Destroy()
		}
	}
}

func createCoinMockSource() domain.CoinSource {
	scrapResultIndex := 0

	coinSource := func(context.Context) <-chan domain.CoinSourceResult {
		coinSourceResultChannel := make(chan domain.CoinSourceResult)

		go func() {
			for _, coin := range scrapResults[scrapResultIndex].Coins {
				coinSourceResultChannel <- domain.CoinSourceResult{
					Coin: coin,
				}
			}

			err := scrapResults[scrapResultIndex].Error
			if err != nil {
				coinSourceResultChannel <- domain.CoinSourceResult{
					Error: err,
				}
			}

			scrapResultIndex += 1

			close(coinSourceResultChannel)
		}()

		return coinSourceResultChannel
	}

	return coinSource
}
