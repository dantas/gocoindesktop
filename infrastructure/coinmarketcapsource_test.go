package infrastructure

import (
	"testing"
)

func TestCoinMarketCapSource(t *testing.T) {
	count := executePrintResultAndReturnCount(t)

	if count == 0 {
		t.Error("Scrapper did not find any coin")
	}
}

func executePrintResultAndReturnCount(t *testing.T) int {
	coins, err := CoinMarketCapSource()

	if err != nil {
		t.Errorf("Errors %v\n", err)
	} else {
		for _, c := range coins {
			t.Logf("Coin %s : %f\n", c.Name, c.Price)
		}
	}

	return len(coins)
}
