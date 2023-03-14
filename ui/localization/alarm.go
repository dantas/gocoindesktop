package localization

import (
	"fmt"

	"github.com/dantas/gocoindesktop/domain/alarm"
)

func AlarmTitle(alarm alarm.TriggeredAlarm) string {
	return fmt.Sprintf("Alert for %s", alarm.Coin.Name)
}

func AlarmEnterRangeMessage(alarm alarm.TriggeredAlarm) string {
	return fmt.Sprintf("Coin %s in range, price: %s", alarm.Coin.Name, FormatPrice(alarm.Coin.Price))
}

func AlarmLeaveRangeMessage(alarm alarm.TriggeredAlarm) string {
	return fmt.Sprintf("Coin %s out of range, price: %s", alarm.Coin.Name, FormatPrice(alarm.Coin.Price))
}
