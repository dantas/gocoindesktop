package presenter

import (
	"github.com/dantas/gocoindesktop/app"
	"github.com/dantas/gocoindesktop/app/alarm"
	"github.com/dantas/gocoindesktop/app/settings"
)

type Event int

const (
	PRESENTER_SHOW_COINS    Event = 0
	PRESENTER_SHOW_SETTINGS Event = iota
)

type Presenter interface {
	OnSystrayClickCoins()
	OnSystrayClickSettings()
	OnSystrayClickQuit()
	SetAlarm(alarm.Alarm)
	Events() <-chan Event
	Settings() settings.Settings
	SetSettings(settings settings.Settings)
	Errors() <-chan error
	CoinAndAlarm() <-chan []app.CoinAndAlarm
	TriggeredAlarms() <-chan alarm.TriggeredAlarm
}
