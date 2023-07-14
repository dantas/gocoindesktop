package fynegui

import (
	"context"

	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"github.com/dantas/gocoindesktop/ui/localization"
	"github.com/getlantern/systray"
)

func Run(presenter Presenter) {
	fyneApp := fyneApp.NewWithID(localization.AppTitle)

	ctx, cancelFunc := context.WithCancel(context.Background())

	go func() {
		<-ctx.Done()
		quit(fyneApp)
	}()

	newWindow(fyneApp, presenter)

	setupSystray(cancelFunc, presenter)

	presenter.Start()

	runMainLoops(fyneApp)
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
