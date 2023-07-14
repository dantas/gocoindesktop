package fynegui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"github.com/dantas/gocoindesktop/ui/localization"
)

type appWindow struct {
	fyne.Window
	app       fyne.App
	presenter Presenter
	tabs      *container.AppTabs
	coinTab   *coinsTab
}

func newWindow(app fyne.App, presenter Presenter) *appWindow {
	window := &appWindow{
		app:       app,
		Window:    app.NewWindow(localization.AppTitle),
		presenter: presenter,
	}

	window.coinTab = newCoinsTab(window, window.presenter)

	window.tabs = container.NewAppTabs(
		container.NewTabItem(localization.TabCoins, window.coinTab),
		container.NewTabItem(localization.TabSettings, newSettingsTab(window.presenter)),
	)

	window.SetContent(window.tabs)
	window.Resize(localization.WindowSize())
	window.SetCloseIntercept(window.Hide)
	window.CenterOnScreen()

	return window
}

func (w *appWindow) Start() {
	go func() {
		for err := range w.presenter.Errors() {
			dialog.ShowError(err, w)
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

			w.app.SendNotification(
				fyne.NewNotification(title, content),
			)
		}
	}()

	go func() {
		for event := range w.presenter.Events() {
			switch event {
			case PRESENTER_SHOW_COINS:
				w.Show()
				w.tabs.SelectIndex(0)
			case PRESENTER_SHOW_SETTINGS:
				w.Show()
				w.tabs.SelectIndex(1)
			}
		}
	}()

	w.coinTab.Start()

	if w.presenter.Settings().ShowWindowOnOpen {
		w.Show()
	}
}
