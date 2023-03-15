//go:build cmd

package ui

import (
	"github.com/dantas/gocoindesktop/app"
	"github.com/dantas/gocoindesktop/ui/cmd"
)

func Run(application *app.Application) {
	cmd.RunCmd(application)
}
