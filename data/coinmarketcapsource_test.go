package data

import (
	"context"
	"testing"
	"time"
)

func TestCoinMarketCapSource(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())

	count := executePrintResultAndReturnCount(ctx, t)

	if count == 0 {
		t.Error("Scrapper did not find any coin")
	}
}

func TestCoinMarketCapSourceCancellation(t *testing.T) {
	firstCtx, _ := context.WithCancel(context.Background())
	firstCount := executePrintResultAndReturnCount(firstCtx, t)

	secondCtx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	secondCount := executePrintResultAndReturnCount(secondCtx, t)

	if firstCount == secondCount {
		t.Errorf("Error stopping second execution, both executions returned same amount of coins, first %d second %d", firstCount, secondCount)
	}
}

func executePrintResultAndReturnCount(ctx context.Context, t *testing.T) uint {
	count := uint(0)

	scrapChannel := CoinMarketCapSource(ctx)

	for result := range scrapChannel {
		if result.Error != nil {
			t.Errorf("Errors %v\n", result.Error)
		} else {
			t.Logf("Coin %s : %f\n", result.Coin.Name, result.Coin.Price)
			count += 1
		}
	}

	return count
}
