package fynegui

import (
	"time"

	"fyne.io/fyne/v2/widget"
	"github.com/dantas/gocoindesktop/domain"
	"github.com/dantas/gocoindesktop/ui/localization"
)

type settingsTab struct {
	*widget.Form
	settings  domain.Settings
	presenter Presenter
}

func newSettingsTab(presenter Presenter) *settingsTab {
	tab := settingsTab{
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

	tab.Form = widget.NewForm(intervalWidget, showWindowOnOpenOption)

	return &tab
}

func (t *settingsTab) newIntervalOption() *widget.Select {
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

		t.settings.Interval = duration
		t.save()
	}

	selectWidget := widget.NewSelect(
		options,
		onSelected,
	)

	switch t.settings.Interval {
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

func (t *settingsTab) newShowWindowOnOpenOption() *widget.Check {
	widget := widget.NewCheck(localization.SettingsShowWindowOnOpenOption, func(isChecked bool) {
		t.settings.ShowWindowOnOpen = isChecked
		t.save()
	})

	widget.Checked = t.settings.ShowWindowOnOpen

	return widget
}

func (t *settingsTab) save() {
	t.presenter.SetSettings(t.settings)
}
