package fynegui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"github.com/dantas/gocoindesktop/ui/localization"
)

type appWindow struct {
	fyne.Window
	tabs      *container.AppTabs
	presenter Presenter
}

func newWindow(app fyne.App, presenter Presenter) *appWindow {
	window := &appWindow{
		Window:    app.NewWindow(localization.AppTitle),
		presenter: presenter,
	}

	go func() {
		for err := range presenter.Errors() {
			dialog.ShowError(err, window)
		}
	}()

	go func() {
		for alarm := range window.presenter.TriggeredAlarms() {
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

	// var appTabs *container.AppTabs

	go func() {
		for event := range window.presenter.Events() {
			switch event {
			case PRESENTER_SHOW_COINS:
				window.Show()
				window.tabs.SelectIndex(0)
			case PRESENTER_SHOW_SETTINGS:
				window.Show()
				window.tabs.SelectIndex(1)
			}
		}
	}()

	window.tabs = container.NewAppTabs(
		container.NewTabItem(localization.TabCoins, newCoinsTab(window, window.presenter)),
		container.NewTabItem(localization.TabSettings, newSettingsTab(window.presenter)),
	)

	window.SetContent(window.tabs)
	window.Resize(localization.WindowSize())
	window.SetCloseIntercept(window.Hide)
	window.CenterOnScreen()

	if presenter.Settings().ShowWindowOnOpen {
		window.Show()
	}

	return window
}
