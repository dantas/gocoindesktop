package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/dantas/gocoindesktop/domain/coin"
	"github.com/dantas/gocoindesktop/ui/localization"
)

func createCoinsTab(window fyne.Window, presenter Presenter) *widget.Table {
	var coins []coin.Coin
	var table *widget.Table

	go func() {
		for coins = range presenter.Coins() {
			table.Refresh()
		}
	}()

	// TODO: Save state in presenter, add timer to reset save timeout

	table = widget.NewTable(
		func() (int, int) {
			return len(coins), 5
		},
		func() fyne.CanvasObject {
			label := widget.NewLabel("")
			label.Alignment = fyne.TextAlignCenter

			return container.NewMax(
				label,
				widget.NewCheck("", nil),
				newNumericalEntry(),
				newNumericalEntry(),
			)
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			label := o.(*fyne.Container).Objects[0].(*widget.Label)
			check := o.(*fyne.Container).Objects[1].(*widget.Check)
			lowerBound := o.(*fyne.Container).Objects[2].(*numericalEntry)
			upperBound := o.(*fyne.Container).Objects[3].(*numericalEntry)

			check.Hide()
			label.Hide()
			lowerBound.Hide()
			upperBound.Hide()

			if i.Row == 0 {
				switch i.Col {
				case 0:
					label.SetText(localization.ColumnCoin)
				case 1:
					label.SetText(localization.ColumnPrice)
				case 2:
					label.SetText(localization.ColumnAlarm)
				case 3:
					label.SetText(localization.ColumnLowerBound)
				case 4:
					label.SetText(localization.ColumnUpperBound)
				}

				label.Show()

				return
			}

			// lowerBound.OnChanged = func(s string) {
			// 	fmt.Println("On Changed")
			// }

			switch i.Col {
			case 0:
				label.SetText(coins[i.Row].Name)
				label.Show()
			case 1:
				label.SetText(localization.FormatPrice(coins[i.Row].Price))
				label.Show()
			case 2:
				check.Show()
			case 3:
				lowerBound.Show()
			case 4:
				upperBound.Show()
			}
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
