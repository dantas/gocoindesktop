package domain

type Application struct {
	timer           PeriodicTimer
	settingsStorage SettingsStorage
	coinSource      CoinSource
	alarmManager    *AlarmManager

	coinsAndAlarms  chan []CoinAndAlarm
	errors          chan error
	triggeredAlarms chan TriggeredAlarm
}

func NewApplication(
	timer PeriodicTimer,
	settingsStorage SettingsStorage,
	coinSource CoinSource,
	alarmManager *AlarmManager,
) *Application {
	return &Application{
		timer:           timer,
		settingsStorage: settingsStorage,
		coinSource:      coinSource,
		alarmManager:    alarmManager,

		coinsAndAlarms:  make(chan []CoinAndAlarm),
		errors:          make(chan error),
		triggeredAlarms: make(chan TriggeredAlarm),
	}
}

func (app *Application) Start() {
	app.timer.SetInterval(app.loadSettings().Interval)

	if err := app.alarmManager.Load(); err != nil {
		app.errors <- err
	}

	go func() {
		app.fetchCoins()

		for range app.timer.Tick() {
			app.fetchCoins()
		}
	}()
}

func (app *Application) CoinsAndAlarms() <-chan []CoinAndAlarm {
	return app.coinsAndAlarms
}

func (app *Application) Errors() <-chan error {
	return app.errors
}

func (app *Application) TriggeredAlarms() <-chan TriggeredAlarm {
	return app.triggeredAlarms
}

func (app *Application) Settings() Settings {
	return app.loadSettings()
}

func (app *Application) SetSettings(settings Settings) {
	err := app.settingsStorage.Save(settings)

	if err == nil {
		app.timer.SetInterval(settings.Interval)
	} else {
		app.errors <- err
	}
}

func (app *Application) SetAlarm(newAlarm Alarm) {
	err := app.alarmManager.Set(newAlarm)

	if err != nil {
		app.errors <- err
	}
}

func (app *Application) Destroy() {
	close(app.coinsAndAlarms)
	close(app.errors)
	close(app.triggeredAlarms)

	app.timer.Destroy()
}

func (app *Application) loadSettings() Settings {
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
