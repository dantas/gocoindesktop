package fynegui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/dantas/gocoindesktop/domain"
	"github.com/dantas/gocoindesktop/ui/localization"
	"github.com/dantas/gocoindesktop/ui/presenter"
)

const _COLUMN_SIZE = 5

func createCoinsTab(window fyne.Window, pres presenter.Presenter) *widget.Table {
	var table *widget.Table

	// Do we need to access this local state through mutex?
	var localCoinAndAlarm []domain.CoinAndAlarm

	go func() {
		for localCoinAndAlarm = range pres.CoinAndAlarm() {
			table.Refresh()
		}
	}()

	table = widget.NewTable(
		func() (int, int) {
			return len(localCoinAndAlarm) + 1, _COLUMN_SIZE
		},
		createRowViews,
		func(i widget.TableCellID, o fyne.CanvasObject) {
			rowViews := getRowViews(o)

			if drawColumnName(i, rowViews) {
				return
			}

			drawContent(i, rowViews, &localCoinAndAlarm[i.Row-1], pres.SetAlarm)
		},
	)

	table.SetColumnWidth(0, localization.ColumnWidthCoin)
	table.SetColumnWidth(1, localization.ColumnWidthPrice)
	table.SetColumnWidth(2, localization.ColumnWidthAlarm)
	table.SetColumnWidth(3, localization.ColumnWidthLowerBound)
	table.SetColumnWidth(4, localization.ColumnWidthUpperBound)

	return table
}

type rowViews struct {
	label      *widget.Label
	check      *widget.Check
	lowerBound *numericalEntry
	upperBound *numericalEntry
}

func createRowViews() fyne.CanvasObject {
	label := widget.NewLabel("")
	label.Alignment = fyne.TextAlignCenter

	return container.NewMax(
		label,
		widget.NewCheck("", nil),
		newNumericalEntry(),
		newNumericalEntry(),
	)
}

func getRowViews(o fyne.CanvasObject) *rowViews {
	rowViews := rowViews{
		label:      o.(*fyne.Container).Objects[0].(*widget.Label),
		check:      o.(*fyne.Container).Objects[1].(*widget.Check),
		lowerBound: o.(*fyne.Container).Objects[2].(*numericalEntry),
		upperBound: o.(*fyne.Container).Objects[3].(*numericalEntry),
	}

	rowViews.label.Hide()
	rowViews.check.Hide()
	rowViews.lowerBound.Hide()
	rowViews.upperBound.Hide()

	return &rowViews
}

func drawColumnName(i widget.TableCellID, rowViews *rowViews) bool {
	if i.Row != 0 {
		return false
	}

	switch i.Col {
	case 0:
		rowViews.label.SetText(localization.ColumnCoin)
	case 1:
		rowViews.label.SetText(localization.ColumnPrice)
	case 2:
		rowViews.label.SetText(localization.ColumnAlarm)
	case 3:
		rowViews.label.SetText(localization.ColumnLowerBound)
	case 4:
		rowViews.label.SetText(localization.ColumnUpperBound)
	}

	rowViews.label.Show()

	return true
}

func drawContent(i widget.TableCellID, rowViews *rowViews, coinAndAlarm *domain.CoinAndAlarm, setAlarm func(domain.Alarm)) {
	switch i.Col {
	case 0:
		rowViews.label.SetText(coinAndAlarm.Coin.Name)
		rowViews.label.Show()
	case 1:
		rowViews.label.SetText(localization.FormatPrice(coinAndAlarm.Coin.Price))
		rowViews.label.Show()
	case 2:
		var checked bool

		if coinAndAlarm.Alarm != nil {
			checked = coinAndAlarm.Alarm.IsEnabled
		}

		rowViews.check.SetChecked(checked)
		rowViews.check.Show()

		rowViews.check.OnChanged = func(isChecked bool) {
			initializeAlarm(coinAndAlarm)
			coinAndAlarm.Alarm.IsEnabled = isChecked
			setAlarm(*coinAndAlarm.Alarm)
		}
	case 3:
		var text string

		if coinAndAlarm.Alarm != nil {
			text = fmt.Sprint(coinAndAlarm.Alarm.LowerBound)
		}

		rowViews.lowerBound.SetText(text)
		rowViews.lowerBound.Show()

		rowViews.lowerBound.OnTextChangedAsFloat64(func(float float64) {
			initializeAlarm(coinAndAlarm)
			coinAndAlarm.Alarm.LowerBound = float
			setAlarm(*coinAndAlarm.Alarm)
		})
	case 4:
		var text string

		if coinAndAlarm.Alarm != nil {
			text = fmt.Sprint(coinAndAlarm.Alarm.UpperBound)
		}

		rowViews.upperBound.SetText(text)
		rowViews.upperBound.Show()

		rowViews.upperBound.OnTextChangedAsFloat64(func(float float64) {
			initializeAlarm(coinAndAlarm)
			coinAndAlarm.Alarm.UpperBound = float
			setAlarm(*coinAndAlarm.Alarm)
		})
	}
}

func initializeAlarm(coinAndAlarm *domain.CoinAndAlarm) {
	if coinAndAlarm.Alarm == nil {
		coinAndAlarm.Alarm = &domain.Alarm{
			Name: coinAndAlarm.Coin.Name,
		}
	}
}
