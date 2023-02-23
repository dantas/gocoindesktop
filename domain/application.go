package domain

type Application struct {
	coinTicker      CoinTicker
	settingsStorage SettingsStorage
	settings        Settings
	errors          chan error
}

func NewApplication(coinTicker CoinTicker, settingsStorage SettingsStorage) Application {
	application := Application{
		coinTicker:      coinTicker,
		settingsStorage: settingsStorage,
		errors:          make(chan error),
	}

	if settings, err := settingsStorage.Load(); err != nil {
		go func() { // TODO: Use goroutine or make channel buffered?
			application.errors <- err
		}()

		application.settings = newDefaultSettings()
	} else {
		application.settings = settings
	}

	coinTicker.SetInterval(application.settings.Interval)

	return application
}

func (app *Application) Coins() <-chan []Coin {
	return app.coinTicker.Coins()
}

func (app *Application) Errors() <-chan error {
	return app.errors
}

func (app *Application) Settings() Settings {
	return app.settings
}

func (app *Application) SetSettings(settings Settings) error {
	if e := app.settingsStorage.Save(settings); e != nil {
		return e
	}

	app.settings = settings
	app.coinTicker.SetInterval(settings.Interval)

	return nil
}

func (app *Application) Destroy() {
	close(app.errors)
	app.coinTicker.Destroy()
}