package domain

type Application struct {
	timer           *periodicTimer
	settingsManager *settingsManager
	coinSource      CoinSource

	coins  chan []Coin
	errors chan error
}

func NewApplication(timer *periodicTimer, settingsManager *settingsManager, coinSource CoinSource) *Application {
	app := Application{
		timer:           timer,
		settingsManager: settingsManager,
		coinSource:      coinSource,

		coins:  make(chan []Coin),
		errors: make(chan error, 1),

		// alarms: make(chan []AlarmTriggered),
		// alarmManager:    AlarmManager{},
	}

	go func() {
		app.fetchCoins()

		for range app.timer.tick {
			// fmt.Println("Fetch")
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
	} else {
		// TODO: Do we need to store coins somewhere in domain?
		app.coins <- coins
	}
}

func (app *Application) Coins() <-chan []Coin {
	return app.coins
}

func (app *Application) Errors() <-chan error {
	return app.errors
}

func (app *Application) Settings() Settings {
	return app.settingsManager.settings
}

func (app *Application) SetSettings(settings Settings) error {
	err := app.settingsManager.SetSettings(settings)

	if err == nil {
		//app.coinTicker.SetInterval(settings.Interval)
		app.timer.SetInterval(settings.Interval)
	}

	return err
}

func (app *Application) Destroy() {
	// close(app.alarms)
	close(app.coins)
	close(app.errors)

	app.timer.Destroy()
}
