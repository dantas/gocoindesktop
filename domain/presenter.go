package domain

type Presenter interface {
	Coins() <-chan []Coin
	Settings() (Settings, error)
	SetSettings(pref Settings) error
	Quit()
}

type presenter struct {
	scrapper        ScrapperTicker
	coins           chan []Coin
	settingsStorage SettingsStorage
}

func NewPresenter(scrapper Scrapper, prefStorage SettingsStorage) Presenter {
	pref, _ := prefStorage.Load()

	app := presenter{
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

func (p presenter) Settings() (Settings, error) {
	return p.settingsStorage.Load()
}

func (p presenter) Coins() <-chan []Coin {
	return p.coins
}

func (p presenter) SetSettings(pref Settings) error {
	if e := p.settingsStorage.Save(pref); e != nil {
		return e
	}
	p.scrapper.SetInterval(pref.Interval)
	return nil
}

func (p presenter) Quit() {
	p.scrapper.Stop()
}
