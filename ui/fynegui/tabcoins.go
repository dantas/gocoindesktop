package fynegui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/dantas/gocoindesktop/domain"
	"github.com/dantas/gocoindesktop/ui/localization"
)

const _COLUMN_SIZE = 5

type coinsTab struct {
	*widget.Table
	coinsAndAlarms []domain.CoinAndAlarm // async code, must prevent lock?
	presenter      Presenter
}

func newCoinsTab(window fyne.Window, presenter Presenter) *coinsTab {
	t := &coinsTab{
		nil,
		make([]domain.CoinAndAlarm, 0),
		presenter,
	}

	t.Table = widget.NewTable(
		t.tableSize,
		createRowViews,
		t.updateCell,
	)

	setColumnWidth(t.Table)

	go func() {
		for t.coinsAndAlarms = range presenter.CoinAndAlarm() {
			t.Table.Refresh()
		}
	}()

	return t
}

func (t *coinsTab) tableSize() (int, int) {
	return len(t.coinsAndAlarms) + 1, _COLUMN_SIZE
}

func (t *coinsTab) updateCell(i widget.TableCellID, o fyne.CanvasObject) {
	rowViews := getRowViews(o)

	if drawColumnName(i, rowViews) {
		return
	}

	t.drawContent(i, rowViews)
}

func (t *coinsTab) drawContent(i widget.TableCellID, rowViews *rowViews) {
	var coinAndAlarm *domain.CoinAndAlarm = &t.coinsAndAlarms[i.Row-1]

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
			t.presenter.SetAlarm(*coinAndAlarm.Alarm)
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
			t.presenter.SetAlarm(*coinAndAlarm.Alarm)
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
			t.presenter.SetAlarm(*coinAndAlarm.Alarm)
		})
	}
}

// Utilitary functions

func initializeAlarm(coinAndAlarm *domain.CoinAndAlarm) {
	if coinAndAlarm.Alarm == nil {
		coinAndAlarm.Alarm = &domain.Alarm{
			Name: coinAndAlarm.Coin.Name,
		}
	}
}

func setColumnWidth(table *widget.Table) {
	table.SetColumnWidth(0, localization.ColumnWidthCoin)
	table.SetColumnWidth(1, localization.ColumnWidthPrice)
	table.SetColumnWidth(2, localization.ColumnWidthAlarm)
	table.SetColumnWidth(3, localization.ColumnWidthLowerBound)
	table.SetColumnWidth(4, localization.ColumnWidthUpperBound)
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

	rowViews.lowerBound.OnTextChangedAsFloat64(func(f float64) {})
	rowViews.upperBound.OnTextChangedAsFloat64(func(f float64) {})

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
