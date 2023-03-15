package gui

import (
	"context"

	"github.com/dantas/gocoindesktop/ui/localization"
	"github.com/getlantern/systray"
)

func setupSystray(cancelFunc context.CancelFunc, presenter Presenter) {
	systray.SetTitle(localization.AppTitle) // app_indicator_set_label: assertion 'IS_APP_INDICATOR (self)' failed
	systray.SetTooltip(localization.AppTitle)
	systray.SetIcon(Icon)

	showCoinsItem := systray.AddMenuItem(localization.SystrayCoins, localization.SystrayCoins)
	showSettingsItem := systray.AddMenuItem(localization.SystraySettings, localization.SystraySettings)
	systray.AddSeparator()
	quitItem := systray.AddMenuItem(localization.SystrayQuit, localization.SystrayQuit)

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
