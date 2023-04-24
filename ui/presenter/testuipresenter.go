package presenter

import (
	"errors"
	"fmt"

	"github.com/dantas/gocoindesktop/domain"
	"github.com/dantas/gocoindesktop/ui/fynegui"
)

type testUi struct {
	events chan fynegui.Event
	errors chan error
}

func NewTestUiPresenter() fynegui.Presenter {
	return &testUi{
		events: make(chan fynegui.Event),
		errors: make(chan error),
	}
}

func (p *testUi) Start() {
	p.errors <- errors.New("something wrong occurred")
}

func (p *testUi) SetAlarm(alarm domain.Alarm) {
	fmt.Printf("Set alarm: %+v\n", alarm)
}

func (p *testUi) OnSystrayClickCoins() {
	p.events <- fynegui.PRESENTER_SHOW_COINS
}

func (p *testUi) OnSystrayClickSettings() {
	p.events <- fynegui.PRESENTER_SHOW_SETTINGS
}

func (p *testUi) OnSystrayClickQuit() {
	close(p.events)
	close(p.errors)
}

func (p *testUi) Events() <-chan fynegui.Event {
	return p.events
}

func (p *testUi) Settings() domain.Settings {
	return domain.NewDefaultSettings()
}

func (p *testUi) SetSettings(settings domain.Settings) {
	fmt.Printf("Set alarm: %+v\n", settings)
}

func (p *testUi) Errors() <-chan error {
	return p.errors
}

func (p *testUi) CoinAndAlarm() <-chan []domain.CoinAndAlarm {
	data := make([]domain.CoinAndAlarm, 0)

	data = append(data, domain.CoinAndAlarm{
		Coin: domain.Coin{
			Name:  "Bitcoin",
			Price: 23400,
		},
		Alarm: &domain.Alarm{
			Name:       "Bitcoin",
			IsEnabled:  false,
			LowerBound: 11000,
			UpperBound: 17000,
		},
	})

	data = append(data, domain.CoinAndAlarm{
		Coin: domain.Coin{
			Name:  "Ethereum",
			Price: 23400,
		},
		Alarm: &domain.Alarm{
			Name:       "Ethereum",
			IsEnabled:  true,
			LowerBound: 2000,
			UpperBound: 5000,
		},
	})

	channel := make(chan []domain.CoinAndAlarm, 1)
	channel <- data
	return channel
}

func (p *testUi) TriggeredAlarms() <-chan domain.TriggeredAlarm {
	data := domain.TriggeredAlarm{
		Alarm: domain.Alarm{
			Name:       "Shiba Inu",
			LowerBound: 1000,
			UpperBound: 2000,
			IsEnabled:  true,
		},
		Coin: domain.Coin{
			Name:  "Shiba Inu",
			Price: 1500,
		},
		InRange: true,
	}

	channel := make(chan domain.TriggeredAlarm, 1)
	channel <- data
	return channel
}
