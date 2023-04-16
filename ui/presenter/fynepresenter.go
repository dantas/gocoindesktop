package presenter

import (
	"github.com/dantas/gocoindesktop/app"
	"github.com/dantas/gocoindesktop/app/alarm"
	"github.com/dantas/gocoindesktop/app/settings"
)

type fynePresenter struct {
	app     *app.Application
	events  chan PresenterEvent
	entries chan []PresenterEntry
}

func NewPresenter(app *app.Application) Presenter {
	presenter := fynePresenter{
		app:     app,
		events:  make(chan PresenterEvent),
		entries: make(chan []PresenterEntry),
	}

	go func() {
		for sliceCoinAlarm := range app.CoinsAndAlarms() {
			presenter.entries <- toSlicePresenterEntry(sliceCoinAlarm)
		}

		close(presenter.entries)
	}()

	return &presenter
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

func (p *fynePresenter) Events() <-chan PresenterEvent {
	return p.events
}

func (p *fynePresenter) Settings() settings.Settings {
	return p.app.Settings()
}

func (p *fynePresenter) SetSettings(settings settings.Settings) {
	p.app.SetSettings(settings)
}

func (p *fynePresenter) Errors() <-chan error {
	return p.app.Errors()
}

func (p *fynePresenter) Entries() <-chan []PresenterEntry {
	return p.entries
}

func (p *fynePresenter) TriggeredAlarms() <-chan alarm.TriggeredAlarm {
	return p.app.TriggeredAlarms()
}

func toSlicePresenterEntry(sliceCoinAlarm []app.CoinAndAlarm) []PresenterEntry {
	entries := make([]PresenterEntry, 0, len(sliceCoinAlarm))

	for _, coinAlarm := range sliceCoinAlarm {
		entry := PresenterEntry{
			Name:  coinAlarm.Coin.Name,
			Price: coinAlarm.Coin.Price,
		}

		if coinAlarm.Alarm != nil {
			entry.IsChecked = true
			entry.LowerBound = coinAlarm.Alarm.LowerBound
			entry.UpperBound = coinAlarm.Alarm.UpperBound
		}

		entries = append(entries, entry)
	}

	return entries
}
