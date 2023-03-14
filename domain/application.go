package domain

import (
	"github.com/dantas/gocoindesktop/domain/alarm"
	"github.com/dantas/gocoindesktop/domain/coin"
)

type CoinSource func() ([]coin.Coin, error)

type Application struct {
	timer           *periodicTimer
	settingsManager *settingsManager
	coinSource      CoinSource
	alarmManager    *alarm.AlarmManager

	coins  chan []coin.Coin
	errors chan error
	alarms chan alarm.TriggeredAlarm
}

func NewApplication(timer *periodicTimer, settingsManager *settingsManager, coinSource CoinSource, alarmManager *alarm.AlarmManager) *Application {
	app := Application{
		timer:           timer,
		settingsManager: settingsManager,
		coinSource:      coinSource,
		alarmManager:    alarmManager,

		coins:  make(chan []coin.Coin),
		errors: make(chan error, 1),
		alarms: make(chan alarm.TriggeredAlarm),
	}

	go func() {
		app.fetchCoins()

		for range app.timer.tick {
			app.fetchCoins()
		}
	}()

	if err := app.settingsManager.Load(); err != nil {
		app.errors <- err
	}

	app.timer.SetInterval(app.settingsManager.settings.Interval)

	return &app
}

func (app *Application) fetchCoins() {
	coins, err := app.coinSource()

	if err != nil {
		app.errors <- err
		return
	}

	app.coins <- coins

	for _, alarm := range app.alarmManager.CheckAlarms(coins) {
		app.alarms <- alarm
	}
}

func (app *Application) Coins() <-chan []coin.Coin {
	return app.coins
}

func (app *Application) Errors() <-chan error {
	return app.errors
}

func (app *Application) Alarms() <-chan alarm.TriggeredAlarm {
	return app.alarms
}

func (app *Application) Settings() Settings {
	return app.settingsManager.settings
}

func (app *Application) SetSettings(settings Settings) error {
	err := app.settingsManager.Save(settings)

	if err == nil {
		app.timer.SetInterval(settings.Interval)
	}

	return err
}

func (app *Application) Destroy() {
	close(app.coins)
	close(app.errors)
	close(app.alarms)

	app.timer.Destroy()
}
