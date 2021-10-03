package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/dantas/gocoindesktop/domain"
	"github.com/dantas/gocoindesktop/ui/localization"
)

func createWindow(app fyne.App, presenter domain.Presenter) fyne.Window {
	window := app.NewWindow(localization.Window.Title)

	window.SetCloseIntercept(func() {
		window.Hide()
	})

	appTabs := container.NewAppTabs(
		container.NewTabItem(localization.Window.TabCoins, createCoinsTab(presenter)),
		container.NewTabItem(localization.Window.TabSetting, createSettingsTab(presenter)),
	)

	window.SetContent(appTabs)

	window.Resize(fyne.NewSize(300, 200))
	window.CenterOnScreen()

	window.Show()

	return window
}
