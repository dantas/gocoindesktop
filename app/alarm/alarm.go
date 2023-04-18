package alarm

import "github.com/dantas/gocoindesktop/app/coin"

type Alarm struct {
	Name       string
	LowerBound float64
	UpperBound float64
	IsEnabled  bool
}

type TriggeredAlarm struct {
	Alarm   Alarm
	Coin    coin.Coin
	InRange bool
}
