package ui

import (
	"github.com/dantas/gocoindesktop/domain"
)

type PresenterShowEvent int

const (
	PRESENTER_SHOW_COINS    = 0
	PRESENTER_SHOW_SETTINGS = iota
)

// TODO: UI/APP Behavior go first to presenter, which will emit events to the UI layer
// Perhaps we would like to mock this in order to test the UI isolated from the rest of the system

type Presenter struct {
	intervalScrapper domain.IntervalScrapper
	settingsStorage  domain.SettingsStorage
	events           chan PresenterShowEvent
}

func NewPresenter(intervalScrapper domain.IntervalScrapper, settingsStorage domain.SettingsStorage) Presenter {
	presenter := Presenter{
		intervalScrapper: intervalScrapper,
		settingsStorage:  settingsStorage,
		events:           make(chan PresenterShowEvent),
	}

	return presenter
}

func (p Presenter) ScrapResults() <-chan domain.ScrapResult {
	return p.intervalScrapper.Results()
}

func (p Presenter) ShowEvents() <-chan PresenterShowEvent {
	return p.events
}

func (p Presenter) OnSystrayClickCoins() {
	p.events <- PRESENTER_SHOW_COINS
}

func (p Presenter) OnSystrayClickSettings() {
	p.events <- PRESENTER_SHOW_SETTINGS
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
