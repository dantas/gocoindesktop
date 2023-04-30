package coinsource

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Execute this test with verbose mode enabled.
// It fetches and prints coins from CoinMarketCap.
func TestCoinMarketCapSource(t *testing.T) {
	coins, err := CoinMarketCapSource()

	assert.Nil(t, err)

	for _, c := range coins {
		t.Logf("Coin %s : %f\n", c.Name, c.Price)
	}

	if len(coins) == 0 {
		t.Error("Scrapper did not find any coin")
	}
}
