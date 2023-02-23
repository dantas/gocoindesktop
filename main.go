package main

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/dantas/gocoindesktop/domain"
	"github.com/dantas/gocoindesktop/infrastructure"
	"github.com/dantas/gocoindesktop/ui"
	"github.com/getlantern/systray"
)

func main() {
	application := newApplicationCompositionRoot()
	runFyneApp(application)
}

func newApplicationCompositionRoot() *domain.Application {
	settingsStorage := infrastructure.NewJsonFileSettingsStorage("settings.json")
	coinTicker := domain.NewCoinTicker(infrastructure.CoinMarketCapSource)
	return domain.NewApplication(coinTicker, settingsStorage)
}

func runFyneApp(application *domain.Application) {
	fyneApp := app.NewWithID("gocoindesktop")

	ctx, cancelFunc := context.WithCancel(context.Background())

	go func() {
		<-ctx.Done()
		quit(fyneApp)
	}()

	setupUi(cancelFunc, fyneApp, application)

	runMainLoops(fyneApp)
}

func setupUi(cancelFunc context.CancelFunc, fyneApp fyne.App, application *domain.Application) {
	presenter := ui.NewPresenter(application)
	ui.CreateWindow(fyneApp, presenter) // TODO: Do we need to keep a reference to the window?
	ui.SetupSystray(cancelFunc, presenter)
}

func runMainLoops(fyneApp fyne.App) {
	// Run systray's main loop in parallel with fyne's main loop
	go func() {
		systray.Run(nil, nil)
	}()

	fyneApp.Run()
}

func quit(fyneApp fyne.App) {
	systray.Quit()
	fyneApp.Quit()
}
