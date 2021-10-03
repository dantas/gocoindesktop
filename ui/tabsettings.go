package ui

import (
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2/widget"
	"github.com/dantas/gocoindesktop/domain"
)

func createSettingsTab(presenter domain.Presenter) *widget.Form {
	settings := presenter.Settings()

	intervalOption := widget.NewFormItem("Update interval", createIntervalOption(&settings))

	form := widget.NewForm(intervalOption)

	form.SubmitText = "Update"
	form.OnSubmit = func() {
		presenter.SetSettings(settings)
	}

	return form
}

func createIntervalOption(settings *domain.Settings) *widget.Select {
	options := []string{"1 min", "2 min", "5 min", "10 min", "1 hour"}

	onSelected := func(selected string) {
		pieces := strings.Split(selected, " ")

		var unit time.Duration
		switch pieces[1] {
		case "min":
			unit = time.Minute
		case "hour":
			unit = time.Hour
		}

		durationInt, _ := strconv.Atoi(pieces[0])

		settings.Interval = time.Duration(durationInt) * unit
	}

	return widget.NewSelect(
		options,
		onSelected,
	)
}
