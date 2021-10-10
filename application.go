package main

import (
	"context"

	"fyne.io/fyne/v2"
	"github.com/dantas/gocoindesktop/ui"
	"github.com/dantas/gocoindesktop/ui/localization"
	"github.com/getlantern/systray"
)

// TODO: Improve name
type Application struct {
	presenter ui.Presenter
	window    fyne.Window
}

func NewApplication(fyneApp fyne.App, presenter ui.Presenter) Application {
	return Application{
		presenter: presenter,
		window:    ui.CreateWindow(fyneApp, presenter),
	}
}

func (uiApp Application) ShowSystray() <-chan struct{} {
	ctx, closeCtx := context.WithCancel(context.Background())

	systray.SetTitle(localization.Window.Title) // app_indicator_set_label: assertion 'IS_APP_INDICATOR (self)' failed
	systray.SetTooltip(localization.Window.Title)
	systray.SetIcon(ui.Icon)

	showCoinsItem := systray.AddMenuItem(localization.Systray.Coins, localization.Systray.Coins)
	showSettingsItem := systray.AddMenuItem(localization.Systray.Settings, localization.Systray.Settings)
	systray.AddSeparator()
	quitItem := systray.AddMenuItem(localization.Systray.Quit, localization.Systray.Quit)

	go func() {
		for {
			select {
			case <-showCoinsItem.ClickedCh:
				uiApp.ShowCoins()
			case <-showSettingsItem.ClickedCh:
				uiApp.ShowSettings()
			case <-quitItem.ClickedCh:
				uiApp.presenter.Quit()
				closeCtx()
			}
		}
	}()

	return ctx.Done()
}

func (app Application) ShowCoins() {
	app.window.Show() // TODO: Redirect to presenter
}

func (app Application) ShowSettings() {
	app.window.Show() // TODO: Redirect to presenter
}
