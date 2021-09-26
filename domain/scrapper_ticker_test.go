package domain_test

import (
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
	scrapperTicker := domain.NewScrapperTicker(createMockScrapper(), 1*time.Second)

	index := 0

	for r := range scrapperTicker.Channel() {
		if len(r) != 1 {
			t.Error("Scrapper returning more values than expected")
		}

		if r[0] != results[index].Coin {
			t.Error("Scrapper didn`t return expected result, index", index)
		}

		index += 1

		switch index {
		case 1:
			index += 1 // Ignore error result
		case 3:
			scrapperTicker.Stop()
		}
	}
}

func createMockScrapper() domain.Scrapper {
	index := 0

	mockScrapper := func(done <-chan interface{}) <-chan domain.ScrapResult {
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
