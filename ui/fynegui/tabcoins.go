package fynegui

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/dantas/gocoindesktop/app"
	"github.com/dantas/gocoindesktop/app/alarm"
	"github.com/dantas/gocoindesktop/ui/localization"
	"github.com/dantas/gocoindesktop/ui/presenter"
)

type rowViews struct {
	label      *widget.Label
	check      *widget.Check
	lowerBound *numericalEntry
	upperBound *numericalEntry
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

const COLUMN_SIZE = 5

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

func drawContent(i widget.TableCellID, rowViews *rowViews, coinAndAlarm *app.CoinAndAlarm, setAlarm func(alarm.Alarm)) {
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

		rowViews.check.OnChanged = func(c bool) {
			if coinAndAlarm.Alarm == nil {
				coinAndAlarm.Alarm = &alarm.Alarm{
					Name: coinAndAlarm.Coin.Name,
				}
			}

			coinAndAlarm.Alarm.IsEnabled = c

			setAlarm(*coinAndAlarm.Alarm)
		}
	case 3:
		var text string

		if coinAndAlarm.Alarm != nil {
			text = fmt.Sprint(coinAndAlarm.Alarm.LowerBound)
		}

		rowViews.lowerBound.SetText(text)
		rowViews.lowerBound.Show()

		rowViews.lowerBound.OnChanged = func(s string) {
			if coinAndAlarm.Alarm == nil {
				coinAndAlarm.Alarm = &alarm.Alarm{
					Name: coinAndAlarm.Coin.Name,
				}
			}

			var number float64
			var err error

			number, err = strconv.ParseFloat(rowViews.lowerBound.Text, 64)

			if err == nil {
				coinAndAlarm.Alarm.LowerBound = number
			}

			setAlarm(*coinAndAlarm.Alarm)
		}
	case 4:
		var text string

		if coinAndAlarm.Alarm != nil {
			text = fmt.Sprint(coinAndAlarm.Alarm.UpperBound)
		}

		rowViews.upperBound.SetText(text)
		rowViews.upperBound.Show()

		rowViews.upperBound.OnChanged = func(s string) {
			if coinAndAlarm.Alarm == nil {
				coinAndAlarm.Alarm = &alarm.Alarm{
					Name: coinAndAlarm.Coin.Name,
				}
			}

			var number float64
			var err error

			number, err = strconv.ParseFloat(rowViews.upperBound.Text, 64)

			if err == nil {
				coinAndAlarm.Alarm.UpperBound = number
			}

			setAlarm(*coinAndAlarm.Alarm)
		}
	}
}

func createCoinsTab(window fyne.Window, pres presenter.Presenter) *widget.Table {
	var table *widget.Table

	// Do we need to access this local state through mutex?
	var localCoinAndAlarm []app.CoinAndAlarm

	go func() {
		for localCoinAndAlarm = range pres.CoinAndAlarm() {
			table.Refresh()
		}
	}()

	table = widget.NewTable(
		func() (int, int) {
			return len(localCoinAndAlarm) + 1, COLUMN_SIZE
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

	// TODO: Move size to another place
	table.SetColumnWidth(0, 150)
	table.SetColumnWidth(1, 100)
	table.SetColumnWidth(2, 100)
	table.SetColumnWidth(3, 100)
	table.SetColumnWidth(4, 100)

	return table
}
