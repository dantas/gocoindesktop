package gui

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/dantas/gocoindesktop/domain"
	"github.com/getlantern/systray"
)

func RunGui(application *domain.Application) {
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
	presenter := newPresenter(application)
	createWindow(fyneApp, presenter) // TODO: Do we need to keep a reference to the window?
	setupSystray(cancelFunc, presenter)
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
