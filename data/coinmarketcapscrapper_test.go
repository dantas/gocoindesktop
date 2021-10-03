package data

import (
	"context"
	"testing"
	"time"
)

func TestCoinMarketCapScrapper(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())

	count := executeScrapPrintResultAndReturnCount(ctx, t)

	if count == 0 {
		t.Error("Scrapper did not find any coin")
	}
}

func TestCoinMarketCapScrapperCancellation(t *testing.T) {
	firstCtx, _ := context.WithCancel(context.Background())
	firstCount := executeScrapPrintResultAndReturnCount(firstCtx, t)

	secondCtx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	secondCount := executeScrapPrintResultAndReturnCount(secondCtx, t)

	if firstCount == secondCount {
		t.Error("Error stopping second execution, both executions returned same ammount of coins")
	}
}

func executeScrapPrintResultAndReturnCount(ctx context.Context, t *testing.T) uint {
	count := uint(0)

	scrapChannel := CoinMarketCapScrapper(ctx)

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
