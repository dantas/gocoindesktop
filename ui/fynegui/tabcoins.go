package fynegui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/dantas/gocoindesktop/ui/localization"
	"github.com/dantas/gocoindesktop/ui/presenter"
)

// type tableState struct {
// 	coins  []coin.Coin
// 	alarms map[string]alarm.Alarm
// }

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

func scheduleSynchronization() {

}

func drawContent(i widget.TableCellID, rowViews *rowViews, entry presenter.Entry) {
	switch i.Col {
	case 0:
		rowViews.label.SetText(entry.Name)
		rowViews.label.Show()
	case 1:
		rowViews.label.SetText(localization.FormatPrice(entry.Price))
		rowViews.label.Show()
	case 2:
		// rowViews.check.OnChanged = func(b bool) {
		// 	scheduleSynchronization()
		// }
		rowViews.check.SetChecked(entry.IsChecked)
		rowViews.check.Show()
	case 3:
		if entry.LowerBound != 0 {
			rowViews.lowerBound.SetText(fmt.Sprint(entry.LowerBound))
		}

		rowViews.lowerBound.Show()
	case 4:
		if entry.UpperBound != 0 {
			rowViews.upperBound.SetText(fmt.Sprint(entry.UpperBound))
		}

		rowViews.upperBound.Show()
	}
}

func createCoinsTab(window fyne.Window, pres presenter.Presenter) *widget.Table {

	// update precisa atualizar coins, sem perder alarmes antigos
	// pegar alarmes salvos e carregar aqui

	var table *widget.Table

	// TODO: Save state in presenter, add timer to reset save timeout
	// Diferenciar entre UI state e alarm state

	var entries []presenter.Entry

	go func() {
		for entries = range pres.Entries() {
			table.Refresh()
		}
	}()

	table = widget.NewTable(
		func() (int, int) {
			return len(entries) + 1, COLUMN_SIZE
		},
		createRowViews,
		func(i widget.TableCellID, o fyne.CanvasObject) {
			rowViews := getRowViews(o)

			if drawColumnName(i, rowViews) {
				return
			}

			drawContent(i, rowViews, entries[i.Row-1])

			// 	case 2:
			// 		_, exist := alarms[coin.Name]

			// 		// Add deboucing
			// 		readNumbersUpdateAlarm := func(row int) {
			// 			alarm, exists := alarms[coin.Name]

			// 			if !exists {
			// 				return
			// 			}

			// 			var number float64
			// 			var err error

			// 			number, err = strconv.ParseFloat(lowerBound.Text, 64)

			// 			if err != nil {
			// 				return
			// 			} else {
			// 				alarm.LowerBound = number
			// 			}

			// 			number, err = strconv.ParseFloat(upperBound.Text, 64)

			// 			if err != nil {
			// 				return
			// 			} else {
			// 				alarm.UpperBound = number
			// 			}
			// 		}

			// 		check.OnChanged = func(isChecked bool) {
			// 			if !isChecked {
			// 				delete(alarms, coin.Name)
			// 				return
			// 			}

			// 			alarms[coin.Name] = alarm.Alarm{
			// 				Name: coin.Name,
			// 			}

			// 			readNumbersUpdateAlarm(i.Row)
			// 		}
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
