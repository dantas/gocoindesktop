package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/dantas/gocoindesktop/data"
	"github.com/dantas/gocoindesktop/domain"
	"github.com/dantas/gocoindesktop/ui"
	"github.com/getlantern/systray"
)

func main() {
	fyneApp := app.New()

	// Our little composition root
	settingsStorage := data.NewFileSettingsStorage("settings.json")
	presenter := domain.NewPresenter(data.CoinMarketCapScrapper, settingsStorage)
	application := ui.NewApplication(fyneApp, presenter)

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
