package fynegui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"github.com/dantas/gocoindesktop/ui/localization"
)

type window struct {
	fyneWindow  fyne.Window
	fyneApp     fyne.App
	fyneTabs    *container.AppTabs
	presenter   Presenter
	tabCoins    *tabCoins
	tabSettings *tabSettings
}

func newWindow(app fyne.App, presenter Presenter) *window {
	window := &window{
		fyneApp:    app,
		fyneWindow: app.NewWindow(localization.AppTitle),
		presenter:  presenter,
	}

	window.tabCoins = newTabCoins(window.fyneWindow, window.presenter)
	window.tabSettings = newTabSettings(window.presenter)

	window.fyneTabs = container.NewAppTabs(
		container.NewTabItem(localization.TabCoins, window.tabCoins.fyneTable),
		container.NewTabItem(localization.TabSettings, window.tabSettings.fyneForm),
	)

	window.fyneWindow.SetContent(window.fyneTabs)
	window.fyneWindow.Resize(localization.WindowSize())
	window.fyneWindow.SetCloseIntercept(window.fyneWindow.Hide)
	window.fyneWindow.CenterOnScreen()

	return window
}

func (w *window) Start() {
	go func() {
		for err := range w.presenter.Errors() {
			dialog.ShowError(err, w.fyneWindow)
		}
	}()

	go func() {
		for alarm := range w.presenter.TriggeredAlarms() {
			title := localization.AlarmTitle(alarm)

			var content string
			if alarm.InRange {
				content = localization.AlarmEnterRangeMessage(alarm)
			} else {
				content = localization.AlarmLeaveRangeMessage(alarm)
			}

			w.fyneApp.SendNotification(
				fyne.NewNotification(title, content),
			)
		}
	}()

	go func() {
		for event := range w.presenter.Events() {
			switch event {
			case PRESENTER_SHOW_COINS:
				w.fyneWindow.Show()
				w.fyneTabs.SelectIndex(0)
			case PRESENTER_SHOW_SETTINGS:
				w.fyneWindow.Show()
				w.fyneTabs.SelectIndex(1)
			}
		}
	}()

	w.tabCoins.Start()

	if w.presenter.Settings().ShowWindowOnOpen {
		w.fyneWindow.Show()
	}
}
