package ui

import (
	"github.com/dantas/gocoindesktop/domain"
)

// TODO: UI/APP Behavior go first to presenter, which will emit events to the UI layer
// Perhaps we would like to mock this in order to test the UI isolated from the rest of the system

type Presenter struct {
	intervalScrapper domain.IntervalScrapper
	settingsStorage  domain.SettingsStorage
}

func NewPresenter(intervalScrapper domain.IntervalScrapper, settingsStorage domain.SettingsStorage) Presenter {
	presenter := Presenter{
		intervalScrapper: intervalScrapper,
		settingsStorage:  settingsStorage,
	}

	return presenter
}

func (p Presenter) ScrapResults() <-chan domain.ScrapResult {
	return p.intervalScrapper.Results()
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
	p.intervalScrapper.Stop()
}
