package domain_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/dantas/gocoindesktop/domain"
)

var results = []domain.ScrapResult{
	{
		Coins: []domain.Coin{
			{
				Name:  "First Coin",
				Price: 12,
			},
			{
				Name:  "Second Coin",
				Price: 42,
			},
			{
				Name:  "Nice Coin",
				Price: 69,
			},
		},
		Errors: []error{},
	},
	{
		Coins: []domain.Coin{},
		Errors: []error{
			fmt.Errorf("first tragic error"),
			fmt.Errorf("second tragic error"),
		},
	},
}

func TestIntervalScrapper(t *testing.T) {
	intervalScrapper := domain.NewIntervalScrapper(createMockScrapper(), 1*time.Second)

	index := 0

	for result := range intervalScrapper.Results() {
		if !reflect.DeepEqual(result, results[index]) {
			t.Errorf("Returned result is different from what is expected %v\n", result)
		}

		index += 1

		if index == len(results) {
			intervalScrapper.Destroy()
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
