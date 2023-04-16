package ui

import (
	"github.com/dantas/gocoindesktop/app"
	"github.com/dantas/gocoindesktop/ui/gui"
)

func Run(application *app.Application) {
	gui.RunGui(application)
}
