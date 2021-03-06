package ui

import (
	"time"

	"github.com/dantas/gocoindesktop/domain"
)

type PresenterShowEvent int

const (
	PRESENTER_SHOW_COINS    = 0
	PRESENTER_SHOW_SETTINGS = iota
)

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

func (p Presenter) UpdateInterval() time.Duration {
	settings, _ := p.settingsStorage.Load()
	return settings.Interval
}

func (p Presenter) SetInterval(interval time.Duration) error {
	settings, _ := p.settingsStorage.Load()
	settings.Interval = interval
	return p.saveSettings(settings)
}

func (p Presenter) ShowWindowOnOpen() bool {
	settings, _ := p.settingsStorage.Load()
	return settings.ShowWindowOnOpen
}

func (p Presenter) SetShowWindowOnOpen(show bool) error {
	settings, _ := p.settingsStorage.Load()
	settings.ShowWindowOnOpen = show
	return p.saveSettings(settings)
}

func (p Presenter) Quit() {
	close(p.events)
	p.intervalScrapper.Destroy()
}

func (p Presenter) saveSettings(settings domain.Settings) error {
	if e := p.settingsStorage.Save(settings); e != nil {
		return e
	}

	p.intervalScrapper.SetInterval(settings.Interval)

	return nil
}
