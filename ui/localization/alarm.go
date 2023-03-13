package localization

import (
	"fmt"

	"github.com/dantas/gocoindesktop/domain/coin"
)

var Alarm struct {
	EnterRange alarmEnterRange
	LeaveRange alarmLeaveRange
}

type alarmEnterRange struct {
	Title string
}

type alarmLeaveRange struct {
	Title string
}

func init() {
	Alarm.EnterRange = alarmEnterRange{
		Title: "Coin entering range",
	}

	Alarm.LeaveRange = alarmLeaveRange{
		Title: "Coin leaving range",
	}
}

func (a alarmEnterRange) Message(coin coin.Coin) string {
	return fmt.Sprintf("Coin: %s - value: %.2f", coin.Name, coin.Price)
}

func (a alarmLeaveRange) Message(coin coin.Coin) string {
	return fmt.Sprintf("Coin: %s - value: %.2f", coin.Name, coin.Price)
}
