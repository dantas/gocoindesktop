package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"github.com/dantas/gocoindesktop/ui/localization"
)

func createWindow(app fyne.App, presenter Presenter) fyne.Window {
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
		for alarm := range presenter.Alarms() {
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
