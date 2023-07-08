package fynegui

import (
	"time"

	"fyne.io/fyne/v2/widget"
	"github.com/dantas/gocoindesktop/ui/localization"
)

func createSettingsTab(pres Presenter) *widget.Form {
	settings := pres.Settings()

	intervalWidget := widget.NewFormItem(
		localization.SettingsUpdateInterval,
		createIntervalOption(settings.Interval, func(interval time.Duration) {
			settings.Interval = interval
			pres.SetSettings(settings)
		}),
	)

	showWindowOnOpenOption := widget.NewFormItem(
		localization.SettingsShowWindowOnOpen,
		createShowWindowOnOpenOption(settings.ShowWindowOnOpen, func(isChecked bool) {
			settings.ShowWindowOnOpen = isChecked
			pres.SetSettings(settings)
		}),
	)

	return widget.NewForm(intervalWidget, showWindowOnOpenOption)
}

func createShowWindowOnOpenOption(initialValue bool, onChanged func(isChecked bool)) *widget.Check {
	widget := widget.NewCheck(localization.SettingsShowWindowOnOpenOption, onChanged)
	widget.Checked = initialValue

	return widget
}

func createIntervalOption(initialValue time.Duration, onChanged func(interval time.Duration)) *widget.Select {
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

		onChanged(duration)
	}

	selectWidget := widget.NewSelect(
		options,
		onSelected,
	)

	switch initialValue {
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
