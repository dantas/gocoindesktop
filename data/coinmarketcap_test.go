package data

import (
	"testing"
)

func TestCoinMarketCap(t *testing.T) {
	done := make(chan interface{})

	scrapChannel := ScrapCoinMarketCap(done)

	for coin := range scrapChannel {
		t.Logf("Coin %s : %f\n", coin.Name, coin.Price)
	}
}
