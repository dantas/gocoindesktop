package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

var window fyne.Window = nil

func main() {
	// var result = scrapper.Scrap()

	// for _, c := range result {
	// 	fmt.Printf("%s : %.2f\n", c.Name, c.Value)
	// }

	application := app.New()

	window = application.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")

	// tabs := container.NewAppTabs()

	// w.SetContent(container.NewVBox(
	// 	hello,
	// 	widget.NewButton("Hi!", func() {
	// 		hello.SetText("Welcome :)")
	// 	}),
	// ))

	// TODO Test without calling this
	application.Driver().ExecuteEveryLoop(systray.RunOnce)

	window.SetContent(hello)

	systray.Register(nil, nil)

	onSystrayReady(application)

	window.ShowAndRun()
}

func onSystrayReady(app fyne.App) {
	systray.AddMenuItem("Settings", "Open settings")

	// systray.SetIcon(icon.Data)
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	// Sets the icon of a menu item. Only available on Mac and Windows.
	mQuit.SetIcon(icon.Data)

	go func() {
		// for i <- mQuit.ClickedCh {

		// }
		for {
			select {
			case <-mQuit.ClickedCh:
				app.Quit()
			}
		}
	}()
}
