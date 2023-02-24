package domain_test

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/dantas/gocoindesktop/domain"
)

var mockCoins = []domain.Coin{
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
}

var errFoo = errors.New("tragic error")

func TestCoinTickerReturnsFetchedCoins(t *testing.T) {
	coinTicker := domain.NewCoinTicker(
		func() ([]domain.Coin, error) {
			return mockCoins, nil
		},
	)

	coinTicker.SetInterval(1 * time.Millisecond)

	for coins := range coinTicker.Coins() {
		if !reflect.DeepEqual(coins, mockCoins) {
			t.Errorf("Returned result is different from what is expected %v != %v\n", coins, mockCoins)
		}

		coinTicker.Destroy()
	}
}

func TestCoinTickerReturnsFailure(t *testing.T) {
	coinTicker := domain.NewCoinTicker(
		func() ([]domain.Coin, error) {
			return make([]domain.Coin, 0), errFoo
		},
	)

	coinTicker.SetInterval(1 * time.Millisecond)

	for err := range coinTicker.Errors() {
		if !errors.Is(err, errFoo) {
			t.Errorf("Returned error is different from what is expected %v != %v\n", err, errFoo)
		}

		coinTicker.Destroy()
	}
}
