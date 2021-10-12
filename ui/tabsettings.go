package ui

import (
	"time"

	"fyne.io/fyne/v2/widget"
	"github.com/dantas/gocoindesktop/ui/localization"
)

func createSettingsTab(presenter Presenter) *widget.Form {
	settings := presenter.Settings()

	intervalOption := widget.NewFormItem(
		localization.Settings.UpdateInterval,
		createIntervalOption(&settings.Interval),
	)

	showWindowOnOpenOption := widget.NewFormItem(
		localization.Settings.ShowWindowOnOpen.FormLabel,
		createShowWindowOnOpenOption(&settings.ShowWindowOnOpen),
	)

	form := widget.NewForm(intervalOption, showWindowOnOpenOption)

	form.SubmitText = localization.Settings.SubmitButton
	form.OnSubmit = func() {
		presenter.SetInterval(settings.Interval)
		presenter.SetShowWindowOnOpen(settings.ShowWindowOnOpen)
	}

	return form
}

func createShowWindowOnOpenOption(show *bool) *widget.Check {
	onCheck := func(isChecked bool) {
		*show = isChecked
	}

	widget := widget.NewCheck(localization.Settings.ShowWindowOnOpen.OptionLabel, onCheck)

	widget.Checked = *show

	return widget
}

func createIntervalOption(interval *time.Duration) *widget.Select {
	options := []string{
		localization.Settings.UpdateIntervalOptions.OneMin,
		localization.Settings.UpdateIntervalOptions.TwoMin,
		localization.Settings.UpdateIntervalOptions.FiveMin,
		localization.Settings.UpdateIntervalOptions.TenMin,
		localization.Settings.UpdateIntervalOptions.OneHour,
	}

	onSelected := func(selected string) {
		switch selected {
		case localization.Settings.UpdateIntervalOptions.OneMin:
			*interval = 1 * time.Minute
		case localization.Settings.UpdateIntervalOptions.TwoMin:
			*interval = 2 * time.Minute
		case localization.Settings.UpdateIntervalOptions.FiveMin:
			*interval = 5 * time.Minute
		case localization.Settings.UpdateIntervalOptions.TenMin:
			*interval = 10 * time.Minute
		case localization.Settings.UpdateIntervalOptions.OneHour:
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
