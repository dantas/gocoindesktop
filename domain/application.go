package domain

type Application struct {
	intervalScrapper IntervalScrapper
	settingsStorage  SettingsStorage
	settings         Settings
}

func NewApplication(intervalScrapper IntervalScrapper, settingsStorage SettingsStorage) Application {
	application := Application{
		intervalScrapper: intervalScrapper,
		settingsStorage:  settingsStorage,
	}

	if settings, err := settingsStorage.Load(); err != nil {
		// TODO: Use exclusive channel for errors?
		application.settings = settings
	} else {
		application.settings = NewDefaultSettings()
	}

	intervalScrapper.SetInterval(application.settings.Interval)

	return application
}

func (app *Application) ScrapResults() <-chan ScrapResult {
	return app.intervalScrapper.Results()
}

func (app Application) Settings() Settings {
	return app.settings
}

func (app Application) SetSettings(settings Settings) error {
	if e := app.settingsStorage.Save(settings); e != nil {
		return e
	}

	app.settings = settings
	app.intervalScrapper.SetInterval(settings.Interval)

	return nil
}

func (app Application) Destroy() {
	app.intervalScrapper.Destroy()
}
