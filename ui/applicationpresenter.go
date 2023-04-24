package ui

import (
	"github.com/dantas/gocoindesktop/domain"
	"github.com/dantas/gocoindesktop/ui/fynegui"
)

type applicationPresenter struct {
	app    *domain.Application
	events chan fynegui.Event
}

func NewApplicationPresenter(app *domain.Application) fynegui.Presenter {
	return &applicationPresenter{
		app:    app,
		events: make(chan fynegui.Event),
	}
}

func (p *applicationPresenter) Start() {
	p.app.Start()
}

func (p *applicationPresenter) SetAlarm(newAlarm domain.Alarm) {
	p.app.SetAlarm(newAlarm)
}

func (p *applicationPresenter) OnSystrayClickCoins() {
	p.events <- fynegui.PRESENTER_SHOW_COINS
}

func (p *applicationPresenter) OnSystrayClickSettings() {
	p.events <- fynegui.PRESENTER_SHOW_SETTINGS
}

func (p *applicationPresenter) OnSystrayClickQuit() {
	close(p.events)
	p.app.Destroy()
}

func (p *applicationPresenter) Events() <-chan fynegui.Event {
	return p.events
}

func (p *applicationPresenter) Settings() domain.Settings {
	return p.app.Settings()
}

func (p *applicationPresenter) SetSettings(settings domain.Settings) {
	p.app.SetSettings(settings)
}

func (p *applicationPresenter) Errors() <-chan error {
	return p.app.Errors()
}

func (p *applicationPresenter) CoinAndAlarm() <-chan []domain.CoinAndAlarm {
	return p.app.CoinsAndAlarms()
}

func (p *applicationPresenter) TriggeredAlarms() <-chan domain.TriggeredAlarm {
	return p.app.TriggeredAlarms()
}
