//go:build cmd

package ui

import (
	"github.com/dantas/gocoindesktop/domain"
	"github.com/dantas/gocoindesktop/ui/cmd"
)

func Run(application *domain.Application) {
	cmd.RunCmd(application)
}
