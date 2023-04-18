package presenter

import (
	"github.com/dantas/gocoindesktop/app"
	"github.com/dantas/gocoindesktop/app/alarm"
	"github.com/dantas/gocoindesktop/app/settings"
)

type fynePresenter struct {
	app     *app.Application
	events  chan Event
	entries chan []Entry
}

func NewPresenter(app *app.Application) Presenter {
	presenter := fynePresenter{
		app:     app,
		events:  make(chan Event),
		entries: make(chan []Entry),
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

func (p *fynePresenter) Events() <-chan Event {
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

func (p *fynePresenter) Entries() <-chan []Entry {
	return p.entries
}

func (p *fynePresenter) TriggeredAlarms() <-chan alarm.TriggeredAlarm {
	return p.app.TriggeredAlarms()
}

func toSlicePresenterEntry(sliceCoinAlarm []app.CoinAndAlarm) []Entry {
	entries := make([]Entry, 0, len(sliceCoinAlarm))

	for _, coinAlarm := range sliceCoinAlarm {
		entry := Entry{
			Name:  coinAlarm.Coin.Name,
			Price: coinAlarm.Coin.Price,
		}

		if coinAlarm.Alarm != nil {
			entry.IsChecked = coinAlarm.Alarm.IsEnabled
			entry.LowerBound = coinAlarm.Alarm.LowerBound
			entry.UpperBound = coinAlarm.Alarm.UpperBound
		}

		entries = append(entries, entry)
	}

	return entries
}
