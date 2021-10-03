package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/dantas/gocoindesktop/domain"
)

func createWindow(app fyne.App, presenter domain.Presenter) fyne.Window {
	window := app.NewWindow("Go Coin Desktop")

	window.SetCloseIntercept(func() {
		window.Hide()
	})

	appTabs := container.NewAppTabs(
		container.NewTabItem("Coins", createCoinsTab(presenter)),
		container.NewTabItem("Settings", createSettingsTab(presenter)),
	)

	window.SetContent(appTabs)

	window.Resize(fyne.NewSize(300, 200))
	window.CenterOnScreen()

	window.Show()

	return window
}
