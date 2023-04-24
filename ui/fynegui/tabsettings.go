package fynegui

import (
	"time"

	"fyne.io/fyne/v2/widget"
	"github.com/dantas/gocoindesktop/domain"
	"github.com/dantas/gocoindesktop/ui/localization"
)

func createSettingsTab(pres Presenter) *widget.Form {
	intervalOption := pres.Settings().Interval
	intervalWidget := widget.NewFormItem(
		localization.SettingsUpdateInterval,
		createIntervalOption(&intervalOption),
	)

	windowOnOpen := pres.Settings().ShowWindowOnOpen
	showWindowOnOpenOption := widget.NewFormItem(
		localization.SettingsShowWindowOnOpen,
		createShowWindowOnOpenOption(&windowOnOpen),
	)

	form := widget.NewForm(intervalWidget, showWindowOnOpenOption)

	form.SubmitText = localization.SettingsSubmitButton
	form.OnSubmit = func() {
		newSettings := domain.Settings{
			Interval:         intervalOption,
			ShowWindowOnOpen: windowOnOpen,
		}

		pres.SetSettings(newSettings)
	}

	return form
}

func createShowWindowOnOpenOption(show *bool) *widget.Check {
	onCheck := func(isChecked bool) {
		*show = isChecked
	}

	widget := widget.NewCheck(localization.SettingsShowWindowOnOpenOption, onCheck)

	widget.Checked = *show

	return widget
}

func createIntervalOption(interval *time.Duration) *widget.Select {
	options := []string{
		localization.Settings1Min,
		localization.Settings2Min,
		localization.Settings5Min,
		localization.Settings10Min,
		localization.Settings1Hour,
	}

	onSelected := func(selected string) {
		switch selected {
		case localization.Settings1Min:
			*interval = 1 * time.Minute
		case localization.Settings2Min:
			*interval = 2 * time.Minute
		case localization.Settings5Min:
			*interval = 5 * time.Minute
		case localization.Settings10Min:
			*interval = 10 * time.Minute
		case localization.Settings1Hour:
			*interval = 1 * time.Hour
		}
	}

	selectWidget := widget.NewSelect(
		options,
		onSelected,
	)

	switch *interval {
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
