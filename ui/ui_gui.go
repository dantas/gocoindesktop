//go:build gui || !cmd

package ui

import (
	"github.com/dantas/gocoindesktop/domain"
	"github.com/dantas/gocoindesktop/ui/gui"
)

func Run(application *domain.Application) {
	gui.RunGui(application)
}
