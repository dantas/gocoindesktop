package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeCoinAndAlarm(t *testing.T) {
	bitcoinCoin := Coin{
		Name:  "Bitcoin",
		Price: 10,
	}

	bitcoinAlarm := Alarm{
		Name:       "Bitcoin",
		LowerBound: 10,
		UpperBound: 15,
		IsEnabled:  true,
	}

	ethereumCoin := Coin{
		Name:  "Ethereum",
		Price: 20,
	}

	ethereumAlarm := Alarm{
		Name:       "Ethereum",
		LowerBound: 20,
		UpperBound: 30,
		IsEnabled:  false,
	}

	shouldBeThereButNoAlarm := Coin{
		Name:  "be_there",
		Price: 12,
	}

	shouldNotBeThere := Alarm{
		Name:       "not_there",
		LowerBound: 23,
		UpperBound: 53,
		IsEnabled:  true,
	}

	sut := merge(
		[]Coin{
			bitcoinCoin, ethereumCoin, shouldBeThereButNoAlarm,
		},
		[]Alarm{
			bitcoinAlarm, ethereumAlarm, shouldNotBeThere,
		},
	)

	assert.Equal(t, len(sut), 3)
	assert.Equal(t, sut[0].Coin, bitcoinCoin)
	assert.Equal(t, *sut[0].Alarm, bitcoinAlarm)
	assert.Equal(t, sut[1].Coin, ethereumCoin)
	assert.Equal(t, *sut[1].Alarm, ethereumAlarm)
	assert.Equal(t, sut[2].Coin, shouldBeThereButNoAlarm)
	assert.Nil(t, sut[2].Alarm)
}
