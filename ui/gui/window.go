package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"github.com/dantas/gocoindesktop/ui/localization"
)

func createWindow(app fyne.App, presenter Presenter) fyne.Window {
	window := app.NewWindow(localization.AppTitle)

	appTabs := container.NewAppTabs(
		container.NewTabItem(localization.TabCoins, createCoinsTab(window, presenter)),
		container.NewTabItem(localization.TabSettings, createSettingsTab(presenter)),
	)

	window.SetContent(appTabs)

	window.Resize(fyne.NewSize(600, 300))
	window.SetCloseIntercept(window.Hide)
	window.CenterOnScreen()

	go func() {
		for event := range presenter.Events() {
			switch event {
			case PRESENTER_SHOW_COINS:
				window.Show()
				appTabs.SelectIndex(0)
			case PRESENTER_SHOW_SETTINGS:
				window.Show()
				appTabs.SelectIndex(1)
			}
		}
	}()

	go func() {
		for err := range presenter.Errors() {
			dialog.ShowError(err, window)
		}
	}()

	go func() {
		for alarm := range presenter.TriggeredAlarms() {
			title := localization.AlarmTitle(alarm)

			var content string
			if alarm.InRange {
				content = localization.AlarmEnterRangeMessage(alarm)
			} else {
				content = localization.AlarmLeaveRangeMessage(alarm)
			}

			app.SendNotification(
				fyne.NewNotification(title, content),
			)
		}
	}()

	if presenter.Settings().ShowWindowOnOpen {
		window.Show()
	}

	return window
}
