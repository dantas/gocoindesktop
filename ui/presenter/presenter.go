package presenter

import (
	"github.com/dantas/gocoindesktop/app/alarm"
	"github.com/dantas/gocoindesktop/app/settings"
)

type Event int

const (
	PRESENTER_SHOW_COINS    Event = 0
	PRESENTER_SHOW_SETTINGS Event = iota
)

type Entry struct {
	Name       string
	Price      float64
	IsChecked  bool
	LowerBound float64
	UpperBound float64
}

type Presenter interface {
	OnSystrayClickCoins()
	OnSystrayClickSettings()
	OnSystrayClickQuit()
	Events() <-chan Event
	Settings() settings.Settings
	SetSettings(settings settings.Settings)
	Errors() <-chan error
	Entries() <-chan []Entry
	TriggeredAlarms() <-chan alarm.TriggeredAlarm
}
