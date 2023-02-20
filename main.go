package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/dantas/gocoindesktop/data"
	"github.com/dantas/gocoindesktop/domain"
	"github.com/dantas/gocoindesktop/ui"
	"github.com/dantas/gocoindesktop/ui/localization"
	"github.com/getlantern/systray"
)

func main() {
	fyneApp := app.NewWithID(localization.App.Title)

	// Our little composition root
	settingsStorage := data.NewJsonFileSettingsStorage("settings.json")
	scrapper := domain.NewScrapper(data.CoinMarketCapSource)
	settings, _ := settingsStorage.Load() // TODO: THIS IS FUCKED UP
	intervalScrapper := domain.NewIntervalScrapper(scrapper, settings.Interval)
	presenter := ui.NewPresenter(intervalScrapper, settingsStorage)
	application := NewApplication(fyneApp, presenter)

	go func() {
		<-application.ShowSystray()
		quit(fyneApp)
	}()

	mainLoop(fyneApp)
}

func mainLoop(fyneApp fyne.App) {
	go func() {
		systray.Run(nil, nil)
	}()

	fyneApp.Run()
}

func quit(fyneApp fyne.App) {
	systray.Quit()
	fyneApp.Quit()
}
