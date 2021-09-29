package domain

type Application struct {
	scrapper        ScrapperTicker
	coins           chan []Coin
	settingsStorage SettingsStorage
}

func NewGoCoinDesktop(scrapper Scrapper, prefStorage SettingsStorage) Application {
	pref, _ := prefStorage.Load()

	app := Application{
		scrapper:        NewScrapperTicker(scrapper, pref.Interval),
		coins:           make(chan []Coin),
		settingsStorage: prefStorage,
	}

	go func() {
		coins := CollectScrapperResults(scrapper(nil))

		if len(coins) > 0 {
			app.coins <- coins
		}
	}()

	go func() {
		for coins := range app.scrapper.Channel() {
			if len(coins) > 0 {
				app.coins <- coins
			}
		}
	}()

	return app
}

func (app *Application) Quit() {
	app.scrapper.Stop()
}

func (app *Application) Coins() <-chan []Coin {
	return app.coins
}

func (app *Application) SetSettings(pref Settings) error {
	if e := app.settingsStorage.Save(pref); e != nil {
		return e
	}
	app.scrapper.SetInterval(pref.Interval)
	return nil
}

func (app *Application) Settings() (Settings, error) {
	return app.settingsStorage.Load()
}
