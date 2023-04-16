package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"github.com/dantas/gocoindesktop/ui/localization"
	"github.com/dantas/gocoindesktop/ui/presenter"
)

func createWindow(app fyne.App, pres presenter.Presenter) fyne.Window {
	window := app.NewWindow(localization.AppTitle)

	appTabs := container.NewAppTabs(
		container.NewTabItem(localization.TabCoins, createCoinsTab(window, pres)),
		container.NewTabItem(localization.TabSettings, createSettingsTab(pres)),
	)

	window.SetContent(appTabs)

	window.Resize(fyne.NewSize(600, 300))
	window.SetCloseIntercept(window.Hide)
	window.CenterOnScreen()

	go func() {
		for event := range pres.Events() {
			switch event {
			case presenter.PRESENTER_SHOW_COINS:
				window.Show()
				appTabs.SelectIndex(0)
			case presenter.PRESENTER_SHOW_SETTINGS:
				window.Show()
				appTabs.SelectIndex(1)
			}
		}
	}()

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

	if pres.Settings().ShowWindowOnOpen {
		window.Show()
	}

	return window
}
