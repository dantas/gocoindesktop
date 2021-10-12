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

	systray.SetTitle(localization.App.Title) // app_indicator_set_label: assertion 'IS_APP_INDICATOR (self)' failed
	systray.SetTooltip(localization.App.Title)
	systray.SetIcon(ui.Icon)

	showCoinsItem := systray.AddMenuItem(localization.Systray.Coins, localization.Systray.Coins)
	showSettingsItem := systray.AddMenuItem(localization.Systray.Settings, localization.Systray.Settings)
	systray.AddSeparator()
	quitItem := systray.AddMenuItem(localization.Systray.Quit, localization.Systray.Quit)

	go func() {
		for {
			select {
			case <-showCoinsItem.ClickedCh:
				uiApp.presenter.OnSystrayClickCoins()
			case <-showSettingsItem.ClickedCh:
				uiApp.presenter.OnSystrayClickSettings()
			case <-quitItem.ClickedCh:
				uiApp.presenter.Quit()
				closeCtx()
			}
		}
	}()

	return ctx.Done()
}
