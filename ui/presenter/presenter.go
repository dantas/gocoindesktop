package presenter

import "github.com/dantas/gocoindesktop/domain"

type Event int

const (
	PRESENTER_SHOW_COINS    Event = 0
	PRESENTER_SHOW_SETTINGS Event = iota
)

type Presenter interface {
	OnSystrayClickCoins()
	OnSystrayClickSettings()
	OnSystrayClickQuit()
	SetAlarm(domain.Alarm)
	Events() <-chan Event
	Settings() domain.Settings
	SetSettings(settings domain.Settings)
	Errors() <-chan error
	CoinAndAlarm() <-chan []domain.CoinAndAlarm
	TriggeredAlarms() <-chan domain.TriggeredAlarm
}
