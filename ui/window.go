package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/dantas/gocoindesktop/ui/localization"
)

func CreateWindow(app fyne.App, presenter Presenter) fyne.Window {
	window := app.NewWindow(localization.Window.Title)

	appTabs := container.NewAppTabs(
		container.NewTabItem(localization.Window.TabCoins, createCoinsTab(presenter)),
		container.NewTabItem(localization.Window.TabSetting, createSettingsTab(presenter)),
	)

	window.SetContent(appTabs)

	window.Resize(fyne.NewSize(300, 200))
	window.SetCloseIntercept(window.Hide)
	window.CenterOnScreen()

	go func() {
		for event := range presenter.ShowEvents() {
			switch event {
			case PRESENTER_SHOW_COINS:
				window.Show()
				appTabs.SelectTabIndex(0)
			case PRESENTER_SHOW_SETTINGS:
				window.Show()
				appTabs.SelectTabIndex(1)
			}
		}
	}()

	// TODO SHOW HIDE DEPENDING ON SETTINGS
	window.Show()

	return window
}
