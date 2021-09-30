package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/dantas/gocoindesktop/domain"
	"github.com/getlantern/systray"
)

type Application struct {
	application domain.Presenter
	window      fyne.Window
}

func (uiApp Application) ShowSystray() <-chan interface{} {
	done := make(chan interface{})

	systray.SetTitle("Go Coin Deskop") // app_indicator_set_label: assertion 'IS_APP_INDICATOR (self)' failed
	systray.SetTooltip("Go Coin Deskop")
	systray.SetIcon(Icon)

	showCoinsItem := systray.AddMenuItem("Show coins", "Show coins")
	showSettingsItem := systray.AddMenuItem("Show settings", "Show settings")
	systray.AddSeparator()
	quitItem := systray.AddMenuItem("Quit", "Quit")

	go func() {
		for {
			select {
			case <-showCoinsItem.ClickedCh:
				uiApp.ShowCoins()
			case <-showSettingsItem.ClickedCh:
				uiApp.ShowSettings()
			case <-quitItem.ClickedCh:
				uiApp.application.Quit()
				close(done)
			}
		}
	}()

	return done
}

func (app Application) ShowCoins() {
	app.window.Show()
}

func (app Application) ShowSettings() {
	app.window.Show()
}

func NewApplication(fyneApp fyne.App, application domain.Presenter) Application {
	return Application{
		application: application,
		window:      createWindow(fyneApp),
	}
}

func createWindow(app fyne.App) fyne.Window {
	window := app.NewWindow("Hello")

	appTabs := container.NewAppTabs(
		container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)

	// Um tab para preferencias
	// Outro tab para grid com resultados, botao de forcar atualizacao

	window.SetContent(appTabs)

	window.CenterOnScreen()

	window.SetCloseIntercept(func() {
		window.Hide()
	})

	window.Hide()

	return window
}
