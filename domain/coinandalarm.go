package domain

type CoinAndAlarm struct {
	Coin  Coin
	Alarm *Alarm
}

func merge(coins []Coin, alarms []Alarm) []CoinAndAlarm {
	result := make([]CoinAndAlarm, 0, len(coins))

	alarmHash := make(map[string]*Alarm)
	for i := range alarms {
		alarmHash[alarms[i].Name] = &alarms[i]
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
