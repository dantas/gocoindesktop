package coin

import (
	"testing"
)

func TestCoinMarketCapSource(t *testing.T) {
	coins, err := CoinMarketCapSource()

	if err != nil {
		t.Errorf("Error scrapping coins: %v\n", err)
	} else {
		for _, c := range coins {
			t.Logf("Coin %s : %f\n", c.Name, c.Price)
		}
	}

	if len(coins) == 0 {
		t.Error("Scrapper did not find any coin")
	}
}
