package domain

type Application struct {
	intervalScrapper IntervalScrapper
	settingsStorage  SettingsStorage
	settings         Settings
	errors           chan error
}

func NewApplication(intervalScrapper IntervalScrapper, settingsStorage SettingsStorage) Application {
	application := Application{
		intervalScrapper: intervalScrapper,
		settingsStorage:  settingsStorage,
		errors:           make(chan error),
	}

	if settings, err := settingsStorage.Load(); err != nil {
		application.errors <- err
		application.settings = newDefaultSettings()
	} else {
		application.settings = settings
	}

	intervalScrapper.SetInterval(application.settings.Interval)

	return application
}

func (app *Application) ScrapResults() <-chan ScrapResult {
	return app.intervalScrapper.Results()
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
	app.intervalScrapper.SetInterval(settings.Interval)

	return nil
}

func (app *Application) Destroy() {
	close(app.errors)
	app.intervalScrapper.Destroy()
}
