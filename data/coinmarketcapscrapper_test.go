package data

import (
	"testing"
	"time"
)

func TestCoinMarketCapScrapper(t *testing.T) {
	done := make(chan interface{})

	count := executeScrapPrintResultAndReturnCount(done, t)

	if count == 0 {
		t.Error("Scrapper did not find any coin")
	}
}

func TestCoinMarketCapScrapperCancellation(t *testing.T) {
	firstDone := make(chan interface{})

	firstCount := executeScrapPrintResultAndReturnCount(firstDone, t)

	secondDone := make(chan interface{})

	time.AfterFunc(1*time.Second, func() {
		close(secondDone)
	})

	secondCount := executeScrapPrintResultAndReturnCount(secondDone, t)

	if firstCount == secondCount {
		t.Error("Error stopping second execution, both executions returned same ammount of coins")
	}
}

func executeScrapPrintResultAndReturnCount(done chan interface{}, t *testing.T) uint {
	count := uint(0)

	scrapChannel := CoinMarketCapScrapper(done)

	for result := range scrapChannel {
		if result.Error != nil {
			t.Error("Error", result.Error)
		} else {
			t.Logf("Coin %s : %f\n", result.Coin.Name, result.Coin.Price)
			count += 1
		}
	}

	return count
}
