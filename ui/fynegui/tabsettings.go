package fynegui

import (
	"time"

	"fyne.io/fyne/v2/widget"
	"github.com/dantas/gocoindesktop/domain"
	"github.com/dantas/gocoindesktop/ui/localization"
)

type tabSettings struct {
	fyneForm  *widget.Form
	settings  domain.Settings
	presenter Presenter
}

func newTabSettings(presenter Presenter) *tabSettings {
	tab := tabSettings{
		nil,
		presenter.Settings(),
		presenter,
	}

	intervalWidget := widget.NewFormItem(
		localization.SettingsUpdateInterval,
		tab.newIntervalOption(),
	)

	showWindowOnOpenOption := widget.NewFormItem(
		localization.SettingsShowWindowOnOpen,
		tab.newShowWindowOnOpenOption(),
	)

	tab.fyneForm = widget.NewForm(intervalWidget, showWindowOnOpenOption)

	return &tab
}

func (tab *tabSettings) newIntervalOption() *widget.Select {
	options := []string{
		localization.Settings1Min,
		localization.Settings2Min,
		localization.Settings5Min,
		localization.Settings10Min,
		localization.Settings1Hour,
	}

	onSelected := func(selected string) {
		var duration time.Duration

		switch selected {
		case localization.Settings1Min:
			duration = 1 * time.Minute
		case localization.Settings2Min:
			duration = 2 * time.Minute
		case localization.Settings5Min:
			duration = 5 * time.Minute
		case localization.Settings10Min:
			duration = 10 * time.Minute
		case localization.Settings1Hour:
			duration = 1 * time.Hour
		}

		tab.settings.Interval = duration
		tab.save()
	}

	selectWidget := widget.NewSelect(
		options,
		onSelected,
	)

	switch tab.settings.Interval {
	case 1 * time.Minute:
		selectWidget.SetSelectedIndex(0)
	case 2 * time.Minute:
		selectWidget.SetSelectedIndex(1)
	case 5 * time.Minute:
		selectWidget.SetSelectedIndex(2)
	case 10 * time.Minute:
		selectWidget.SetSelectedIndex(3)
	case 1 * time.Hour:
		selectWidget.SetSelectedIndex(4)
	}

	return selectWidget
}

func (tab *tabSettings) newShowWindowOnOpenOption() *widget.Check {
	widget := widget.NewCheck(localization.SettingsShowWindowOnOpenOption, func(isChecked bool) {
		tab.settings.ShowWindowOnOpen = isChecked
		tab.save()
	})

	widget.Checked = tab.settings.ShowWindowOnOpen

	return widget
}

func (tab *tabSettings) save() {
	tab.presenter.SetSettings(tab.settings)
}
