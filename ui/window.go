package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/dantas/gocoindesktop/ui/localization"
)

func CreateWindow(app fyne.App, presenter Presenter) fyne.Window {
	window := app.NewWindow(localization.App.Title)

	appTabs := container.NewAppTabs(
		container.NewTabItem(localization.Window.TabCoins, createCoinsTab(window, presenter)),
		container.NewTabItem(localization.Window.TabSetting, createSettingsTab(presenter)),
	)

	window.SetContent(appTabs)

	window.Resize(fyne.NewSize(400, 300))
	window.SetCloseIntercept(window.Hide)
	window.CenterOnScreen()

	go func() {
		for event := range presenter.Events() {
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

	// go func() {
	// 	for event := range presenter.AlarmEvents() {
	// 		notification := fyne.NewNotification(
	// 			localization.Alarm.EnterRange.Title,
	// 			localization.Alarm.EnterRange.Message(event.Coin),
	// 		)

	// 		app.SendNotification(notification)
	// 	}
	// }()

	if presenter.Settings().ShowWindowOnOpen {
		window.Show()
	}

	return window
}
