package main

import (
	"github.com/dantas/gocoindesktop/compositionroot"
	"github.com/dantas/gocoindesktop/ui/fynegui"
)

func main() {
	presenter := compositionroot.NewPresenter()
	fynegui.Run(presenter)
}
