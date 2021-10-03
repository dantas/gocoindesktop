package ui

import (
	"time"

	"fyne.io/fyne/v2/widget"
	"github.com/dantas/gocoindesktop/domain"
	"github.com/dantas/gocoindesktop/ui/localization"
)

func createSettingsTab(presenter domain.Presenter) *widget.Form {
	settings := presenter.Settings()

	intervalOption := widget.NewFormItem(localization.Settings.UpdateInterval, createIntervalOption(&settings))

	form := widget.NewForm(intervalOption)

	form.SubmitText = localization.Settings.SubmitButton
	form.OnSubmit = func() {
		presenter.SetSettings(settings)
	}

	return form
}

func createIntervalOption(settings *domain.Settings) *widget.Select {
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
			settings.Interval = 1 * time.Minute
		case localization.Settings.UpdateIntervalOptions.TwoMin:
			settings.Interval = 2 * time.Minute
		case localization.Settings.UpdateIntervalOptions.FiveMin:
			settings.Interval = 5 * time.Minute
		case localization.Settings.UpdateIntervalOptions.TenMin:
			settings.Interval = 10 * time.Minute
		case localization.Settings.UpdateIntervalOptions.OneHour:
			settings.Interval = 1 * time.Hour
		}
	}

	selectWidget := widget.NewSelect(
		options,
		onSelected,
	)

	switch settings.Interval {
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
