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

func createIntervalOption(settings *domain.Settings) *widget.RadioGroup {
	options := []string{"1 min", "2 min", "5 min", "10 min", "1 hour"}

	onSelected := func(selected string) {
		durationStr := strings.TrimSuffix(selected, " min")
		durationInt, _ := strconv.Atoi(durationStr)
		settings.Interval = time.Duration(durationInt) * time.Minute
	}

	return widget.NewRadioGroup(
		options,
		onSelected,
	)
}
