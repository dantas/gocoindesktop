//go:build testui

package compositionroot

import (
	"github.com/dantas/gocoindesktop/ui/fynegui"
	"github.com/dantas/gocoindesktop/ui/presenter"
)

func NewPresenter() fynegui.Presenter {
	return presenter.NewTestUiPresenter()
}
