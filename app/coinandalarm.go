package app

import (
	"github.com/dantas/gocoindesktop/app/alarm"
	"github.com/dantas/gocoindesktop/app/coin"
)

type CoinAndAlarm struct {
	Coin  coin.Coin
	Alarm *alarm.Alarm
}

func merge(coins []coin.Coin, alarms []alarm.Alarm) []CoinAndAlarm {
	result := make([]CoinAndAlarm, 0, len(coins))

	alarmHash := make(map[string]*alarm.Alarm)
	for _, alarm := range alarms {
		alarmHash[alarm.Name] = &alarm
	}

	for _, coin := range coins {
		result = append(
			result,
			CoinAndAlarm{
				Coin:  coin,
				Alarm: alarmHash[coin.Name],
			},
		)
	}

	return result
}
