package gui

import (
	"context"

	"github.com/dantas/gocoindesktop/ui/localization"
	"github.com/getlantern/systray"
)

func SetupSystray(cancelFunc context.CancelFunc, presenter Presenter) {
	systray.SetTitle(localization.App.Title) // app_indicator_set_label: assertion 'IS_APP_INDICATOR (self)' failed
	systray.SetTooltip(localization.App.Title)
	systray.SetIcon(Icon)

	showCoinsItem := systray.AddMenuItem(localization.Systray.Coins, localization.Systray.Coins)
	showSettingsItem := systray.AddMenuItem(localization.Systray.Settings, localization.Systray.Settings)
	systray.AddSeparator()
	quitItem := systray.AddMenuItem(localization.Systray.Quit, localization.Systray.Quit)

	go func() {
		for {
			select {
			case <-showCoinsItem.ClickedCh:
				presenter.OnSystrayClickCoins()
			case <-showSettingsItem.ClickedCh:
				presenter.OnSystrayClickSettings()
			case <-quitItem.ClickedCh:
				presenter.OnSystrayClickQuit()
				cancelFunc()
			}
		}
	}()
}
