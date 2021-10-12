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

// TODO Move to domain
type PresenterAlarmEvent struct {
	Coin domain.Coin
}

// Perhaps we would like to mock this in order to test the UI isolated from the rest of the system

type Presenter struct {
	intervalScrapper domain.IntervalScrapper
	settingsStorage  domain.SettingsStorage
	showEvents       chan PresenterShowEvent
	alarmEvents      chan PresenterAlarmEvent
}

func NewPresenter(intervalScrapper domain.IntervalScrapper, settingsStorage domain.SettingsStorage) Presenter {
	presenter := Presenter{
		intervalScrapper: intervalScrapper,
		settingsStorage:  settingsStorage,
		showEvents:       make(chan PresenterShowEvent),
		alarmEvents:      make(chan PresenterAlarmEvent),
	}

	return presenter
}

func (p Presenter) ScrapResults() <-chan domain.ScrapResult {
	return p.intervalScrapper.Results()
}

func (p Presenter) ShowEvents() <-chan PresenterShowEvent {
	return p.showEvents
}

func (p Presenter) AlarmEvents() <-chan PresenterAlarmEvent {
	return p.alarmEvents
}

func (p Presenter) OnSystrayClickCoins() {
	p.showEvents <- PRESENTER_SHOW_COINS
}

func (p Presenter) OnSystrayClickSettings() {
	p.showEvents <- PRESENTER_SHOW_SETTINGS
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
	close(p.showEvents)
	close(p.alarmEvents)
	p.intervalScrapper.Destroy()
}

func (p Presenter) saveSettings(settings domain.Settings) error {
	if e := p.settingsStorage.Save(settings); e != nil {
		return e
	}

	p.intervalScrapper.SetInterval(settings.Interval)

	return nil
}
