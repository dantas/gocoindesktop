package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/dantas/gocoindesktop/domain"
	"github.com/dantas/gocoindesktop/infrastructure"
	"github.com/dantas/gocoindesktop/ui"
	"github.com/getlantern/systray"
)

func StartApplication(fyneApp fyne.App) <-chan struct{} {
	// Our little composition root
	settingsStorage := infrastructure.NewJsonFileSettingsStorage("settings.json")
	scrapper := domain.NewScrapper(infrastructure.CoinMarketCapSource)
	intervalScrapper := domain.NewIntervalScrapper(scrapper)

	application := domain.NewApplication(intervalScrapper, settingsStorage)

	presenter := ui.NewPresenter(application)

	ui.CreateWindow(fyneApp, presenter) // TODO: Do we need to keep a reference to this?

	return ui.CreateSystray(presenter)
}

func main() {
	fyneApp := app.NewWithID("gocoindesktop") // TODO Double check id

	go func() {
		<-StartApplication(fyneApp)
		quit(fyneApp)
	}()

	mainLoop(fyneApp)
}

func mainLoop(fyneApp fyne.App) {
	// Connect systray main loop with fyne main loop
	go func() {
		systray.Run(nil, nil)
	}()

	fyneApp.Run()
}

func quit(fyneApp fyne.App) {
	systray.Quit()
	fyneApp.Quit()
}
