package fynegui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"github.com/dantas/gocoindesktop/ui/localization"
)

func createWindow(app fyne.App, pres Presenter) fyne.Window {
	var window fyne.Window

	go func() {
		for err := range pres.Errors() {
			dialog.ShowError(err, window)
		}
	}()

	go func() {
		for alarm := range pres.TriggeredAlarms() {
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

	var appTabs *container.AppTabs

	go func() {
		for event := range pres.Events() {
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

	window = app.NewWindow(localization.AppTitle)

	appTabs = container.NewAppTabs(
		container.NewTabItem(localization.TabCoins, newCoinsTab(window, pres)),
		container.NewTabItem(localization.TabSettings, newSettingsTab(pres)),
	)

	window.SetContent(appTabs)
	window.Resize(localization.WindowSize())
	window.SetCloseIntercept(window.Hide)
	window.CenterOnScreen()

	if pres.Settings().ShowWindowOnOpen {
		window.Show()
	}

	return window
}
