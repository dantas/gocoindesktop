package presenter

import "github.com/dantas/gocoindesktop/domain"

type fynePresenter struct {
	app    *domain.Application
	events chan Event
}

func NewPresenter(app *domain.Application) Presenter {
	presenter := fynePresenter{
		app:    app,
		events: make(chan Event),
	}

	return &presenter
}

func (p *fynePresenter) SetAlarm(newAlarm domain.Alarm) {
	p.app.SetAlarm(newAlarm)
}

func (p *fynePresenter) OnSystrayClickCoins() {
	p.events <- PRESENTER_SHOW_COINS
}

func (p *fynePresenter) OnSystrayClickSettings() {
	p.events <- PRESENTER_SHOW_SETTINGS
}

func (p *fynePresenter) OnSystrayClickQuit() {
	close(p.events)
	p.app.Destroy()
}

func (p *fynePresenter) Events() <-chan Event {
	return p.events
}

func (p *fynePresenter) Settings() domain.Settings {
	return p.app.Settings()
}

func (p *fynePresenter) SetSettings(settings domain.Settings) {
	p.app.SetSettings(settings)
}

func (p *fynePresenter) Errors() <-chan error {
	return p.app.Errors()
}

func (p *fynePresenter) CoinAndAlarm() <-chan []domain.CoinAndAlarm {
	return p.app.CoinsAndAlarms()
}

func (p *fynePresenter) TriggeredAlarms() <-chan domain.TriggeredAlarm {
	return p.app.TriggeredAlarms()
}
