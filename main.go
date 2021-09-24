package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/dantas/gocoindesktop/ui"
	"github.com/getlantern/systray"
)

func main() {
	app := app.New()

	configureSystray(app)

	mainLoop(app)
}

func configureSystray(app fyne.App) {
	systray.SetTitle("Go Coin Deskop") // app_indicator_set_label: assertion 'IS_APP_INDICATOR (self)' failed
	systray.SetTooltip("Go Coin Deskop")
	systray.SetIcon(ui.Icon)

	openApplicationItem := systray.AddMenuItem("Open application", "Open application")
	openSettingsItem := systray.AddMenuItem("Open settings", "Open settings")
	systray.AddSeparator()
	quitItem := systray.AddMenuItem("Quit", "Quit")

	go func() {
		for {
			select {
			case <-openApplicationItem.ClickedCh:
				ui.OpenApplication(app)
			case <-openSettingsItem.ClickedCh:
				ui.OpenSettings()
			case <-quitItem.ClickedCh:
				quit(app)
			}
		}
	}()
}

func mainLoop(app fyne.App) {
	go func() {
		systray.Run(nil, nil)
	}()

	app.Run()
}

func quit(app fyne.App) {
	systray.Quit()
	app.Quit()
}
