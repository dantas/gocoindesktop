package alarm

import "github.com/dantas/gocoindesktop/domain/coin"

type Alarm struct {
	Name       string
	LowerBound float64
	UpperBound float64
}

type TriggeredAlarm struct {
	Alarm   Alarm
	Coin    coin.Coin
	InRange bool
}
