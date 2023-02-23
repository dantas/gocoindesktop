package ui

import (
	"github.com/dantas/gocoindesktop/domain"
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
	Settings() domain.Settings
	SetSettings(settings domain.Settings)
	Coins() <-chan []domain.Coin
	Errors() <-chan error
}

type presenter struct {
	app    *domain.Application
	events chan PresenterEvent
}

func NewPresenter(app *domain.Application) Presenter {
	return presenter{
		app:    app,
		events: make(chan PresenterEvent),
	}
}

func (p presenter) OnSystrayClickCoins() {
	p.events <- PRESENTER_SHOW_COINS
}

func (p presenter) OnSystrayClickSettings() {
	p.events <- PRESENTER_SHOW_SETTINGS
}

func (p presenter) OnSystrayClickQuit() {
	close(p.events)
	p.app.Destroy()
}

func (p presenter) Events() <-chan PresenterEvent {
	return p.events
}

func (p presenter) Settings() domain.Settings {
	return p.app.Settings()
}

func (p presenter) SetSettings(settings domain.Settings) {
	p.app.SetSettings(settings)
}

func (p presenter) Coins() <-chan []domain.Coin {
	return p.app.Coins()
}

func (p presenter) Errors() <-chan error {
	return p.app.Errors()
}
