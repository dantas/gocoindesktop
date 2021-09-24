package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var window fyne.Window = nil

// Keep reference to this windows somewhere
func OpenApplication(app fyne.App) {
	if window == nil {

		window = app.NewWindow("Hello")

		window.CenterOnScreen()

		window.SetCloseIntercept(func() {
			window.Hide()
		})
	}

	appTabs := container.NewAppTabs(
		container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)

	// Um tab para preferencias
	// Outro tab para grid com resultados, botao de forcar atualizacao

	window.SetContent(appTabs)

	window.Show()
	// window.Hide()
}
