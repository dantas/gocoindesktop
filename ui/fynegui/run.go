package fynegui

import (
	"context"

	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"github.com/dantas/gocoindesktop/domain"
	"github.com/dantas/gocoindesktop/ui/localization"
	"github.com/dantas/gocoindesktop/ui/presenter"
	"github.com/getlantern/systray"
)

func Run(application *domain.Application) {
	fyneApp := fyneApp.NewWithID(localization.AppTitle)

	ctx, cancelFunc := context.WithCancel(context.Background())

	go func() {
		<-ctx.Done()
		quit(fyneApp)
	}()

	setupUi(cancelFunc, fyneApp, application)

	application.Start()

	runMainLoops(fyneApp)
}

func setupUi(cancelFunc context.CancelFunc, fyneApp fyne.App, application *domain.Application) {
	presenter := presenter.NewPresenter(application)
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
