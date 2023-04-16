package presenter

import (
	"github.com/dantas/gocoindesktop/app/alarm"
	"github.com/dantas/gocoindesktop/app/settings"
)

type PresenterEvent int

const (
	PRESENTER_SHOW_COINS    PresenterEvent = 0
	PRESENTER_SHOW_SETTINGS PresenterEvent = iota
)

// Find a better name?
type PresenterEntry struct {
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
	Events() <-chan PresenterEvent
	Settings() settings.Settings
	SetSettings(settings settings.Settings)
	Errors() <-chan error
	Entries() <-chan []PresenterEntry
	TriggeredAlarms() <-chan alarm.TriggeredAlarm
}
