package ui

import (
	"github.com/dantas/gocoindesktop/domain"
)

type PresenterEvent int

const (
	EVENT_SHOW_COINS    = 0
	EVENT_SHOW_SETTINGS = iota
)

// TODO: UI/APP Behavior go first to presenter, which will emit events to the UI layer
// Perhaps we would like to mock this in order to test the UI isolated from the rest of the system

type Presenter struct {
	intervalScrapper domain.IntervalScrapper
	settingsStorage  domain.SettingsStorage
	events           chan PresenterEvent
}

func NewPresenter(intervalScrapper domain.IntervalScrapper, settingsStorage domain.SettingsStorage) Presenter {
	presenter := Presenter{
		intervalScrapper: intervalScrapper,
		settingsStorage:  settingsStorage,
		events:           make(chan PresenterEvent),
	}

	return presenter
}

func (p Presenter) ScrapResults() <-chan domain.ScrapResult {
	return p.intervalScrapper.Results()
}

func (p Presenter) Events() <-chan PresenterEvent {
	return p.events
}

func (p Presenter) OnSystrayClickCoins() {
	p.events <- EVENT_SHOW_COINS
}

func (p Presenter) OnSystrayClickSettings() {
	p.events <- EVENT_SHOW_SETTINGS
}

func (p Presenter) Settings() domain.Settings {
	settings, _ := p.settingsStorage.Load()
	return settings
}

func (p Presenter) SetSettings(settings domain.Settings) error {
	if e := p.settingsStorage.Save(settings); e != nil {
		return e
	}

	p.intervalScrapper.SetInterval(settings.Interval)

	return nil
}

func (p Presenter) Quit() {
	close(p.events)
	p.intervalScrapper.Stop()
}
