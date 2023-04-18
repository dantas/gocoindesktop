package app

import (
	"github.com/dantas/gocoindesktop/app/alarm"
	"github.com/dantas/gocoindesktop/app/coin"
	"github.com/dantas/gocoindesktop/app/settings"
	"github.com/dantas/gocoindesktop/app/timer"
)

type Application struct {
	timer           *timer.PeriodicTimer
	settingsStorage settings.SettingsStorage
	coinSource      coin.CoinSource
	alarmManager    *alarm.AlarmManager

	coinsAndAlarms  chan []CoinAndAlarm
	errors          chan error
	triggeredAlarms chan alarm.TriggeredAlarm
}

func NewApplication(
	timer *timer.PeriodicTimer,
	settingsStorage settings.SettingsStorage,
	coinSource coin.CoinSource,
	alarmManager *alarm.AlarmManager,
) *Application {
	app := Application{
		timer:           timer,
		settingsStorage: settingsStorage,
		coinSource:      coinSource,
		alarmManager:    alarmManager,

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

	app.timer.SetInterval(app.loadSettings().Interval)

	if err := app.alarmManager.Load(); err != nil {
		app.errors <- err
	}

	return &app
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
	return app.loadSettings()
}

func (app *Application) SetSettings(settings settings.Settings) {
	err := app.settingsStorage.Save(settings)

	if err == nil {
		app.timer.SetInterval(settings.Interval)
	} else {
		app.errors <- err
	}
}

func (app *Application) Destroy() {
	close(app.coinsAndAlarms)
	close(app.errors)
	close(app.triggeredAlarms)

	app.timer.Destroy()
}

func (app *Application) loadSettings() settings.Settings {
	sett, err := app.settingsStorage.Load()

	if err != nil {
		app.errors <- err
	}

	return sett
}

func (app *Application) fetchCoins() {
	coins, err := app.coinSource()

	if err != nil {
		app.errors <- err
		return
	}

	app.coinsAndAlarms <- merge(coins, app.alarmManager.Alarms())

	for _, alarm := range app.alarmManager.CheckAlarms(coins) {
		app.triggeredAlarms <- alarm
	}
}
