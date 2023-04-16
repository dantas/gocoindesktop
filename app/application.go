package app

import (
	"github.com/dantas/gocoindesktop/app/alarm"
	"github.com/dantas/gocoindesktop/app/coin"
	"github.com/dantas/gocoindesktop/app/settings"
	"github.com/dantas/gocoindesktop/app/timer"
)

type Application struct {
	timer              *timer.PeriodicTimer
	settingsRepository *settings.Repository
	coinSource         coin.CoinSource
	alarmManager       *alarm.AlarmManager

	coinsAndAlarms  chan []CoinAndAlarm
	errors          chan error
	triggeredAlarms chan alarm.TriggeredAlarm
}

func NewApplication(
	timer *timer.PeriodicTimer,
	settingsRepository *settings.Repository,
	coinSource coin.CoinSource,
	alarmManager *alarm.AlarmManager,
) *Application {
	app := Application{
		timer:              timer,
		settingsRepository: settingsRepository,
		coinSource:         coinSource,
		alarmManager:       alarmManager,

		coinsAndAlarms:  make(chan []CoinAndAlarm),
		errors:          make(chan error, 1),
		triggeredAlarms: make(chan alarm.TriggeredAlarm),
	}

	go func() {
		app.fetchCoins()

		for range app.timer.Tick() {
			app.fetchCoins()
		}
	}()

	settings, err := settingsRepository.Load()

	if err != nil {
		app.errors <- err
	}

	app.timer.SetInterval(settings.Interval)

	return &app
}

func (app *Application) fetchCoins() {
	coins, err := app.coinSource()

	if err != nil {
		app.errors <- err
		return
	}

	// TODO: Load alarms from alarm manager
	app.coinsAndAlarms <- merge(coins, nil)

	for _, alarm := range app.alarmManager.CheckAlarms(coins) {
		app.triggeredAlarms <- alarm
	}
}

func (app *Application) CoinsAndAlarms() <-chan []CoinAndAlarm {
	return app.coinsAndAlarms
}

func (app *Application) Errors() <-chan error {
	return app.errors
}

func (app *Application) TriggeredAlarms() <-chan alarm.TriggeredAlarm {
	return app.triggeredAlarms
}

func (app *Application) Settings() settings.Settings {
	settings, _ := app.settingsRepository.Load() // TODO: Error handling?
	return settings
}

func (app *Application) SetSettings(settings settings.Settings) error {
	err := app.settingsRepository.Save(settings)

	if err == nil {
		app.timer.SetInterval(settings.Interval)
	}

	// Return error or notify channel of error?
	return err
}

func (app *Application) Destroy() {
	close(app.coinsAndAlarms)
	close(app.errors)
	close(app.triggeredAlarms)

	app.timer.Destroy()
}
