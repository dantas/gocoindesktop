package domain

type CoinAndAlarm struct {
	Coin  Coin
	Alarm *Alarm
}

func merge(coins []Coin, alarms []Alarm) []CoinAndAlarm {
	result := make([]CoinAndAlarm, 0, len(coins))

	alarmHash := make(map[string]*Alarm)
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
