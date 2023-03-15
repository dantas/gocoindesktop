package gui

import (
	"github.com/dantas/gocoindesktop/app/alarm"
	"github.com/dantas/gocoindesktop/app/settings"
	"github.com/dantas/gocoindesktop/domain"
	"github.com/dantas/gocoindesktop/domain/coin"
)

type PresenterEvent int

const (
	PRESENTER_SHOW_COINS    PresenterEvent = 0
	PRESENTER_SHOW_SETTINGS PresenterEvent = iota
)

type Presenter interface {
	OnSystrayClickCoins()
	OnSystrayClickSettings()
	OnSystrayClickQuit()
	Events() <-chan PresenterEvent
	Settings() settings.Settings
	SetSettings(settings settings.Settings)
	Coins() <-chan []coin.Coin
	Errors() <-chan error
	Alarms() <-chan alarm.TriggeredAlarm
}

type presenter struct {
	app    *domain.Application
	events chan PresenterEvent
}

func newPresenter(app *domain.Application) Presenter {
	return &presenter{
		app:    app,
		events: make(chan PresenterEvent),
	}
}

func (p *presenter) OnSystrayClickCoins() {
	p.events <- PRESENTER_SHOW_COINS
}

func (p *presenter) OnSystrayClickSettings() {
	p.events <- PRESENTER_SHOW_SETTINGS
}

func (p *presenter) OnSystrayClickQuit() {
	close(p.events)
	p.app.Destroy()
}

func (p *presenter) Events() <-chan PresenterEvent {
	return p.events
}

func (p *presenter) Settings() settings.Settings {
	return p.app.Settings()
}

func (p *presenter) SetSettings(settings settings.Settings) {
	p.app.SetSettings(settings)
}

func (p *presenter) Coins() <-chan []coin.Coin {
	return p.app.Coins()
}

func (p *presenter) Errors() <-chan error {
	return p.app.Errors()
}

func (p *presenter) Alarms() <-chan alarm.TriggeredAlarm {
	return p.app.Alarms()
}
