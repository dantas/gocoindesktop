package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/dantas/gocoindesktop/domain"
)

func createCoinsTab(presenter Presenter) *widget.Table {
	var coins []domain.Coin
	var table *widget.Table

	var column [2]float32

	go func() {
		for result := range presenter.ScrapResults() {
			coins = result.Coins
			// TODO: Show errors (using dialogs) when we receive them from the domain layer
			table.Refresh()
			table.SetColumnWidth(0, column[0])
			table.SetColumnWidth(1, column[1])
		}
	}()

	table = widget.NewTable(
		func() (int, int) {
			return len(coins), 2
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			label := o.(*widget.Label)
			var content string
			switch i.Col {
			case 0:
				content = coins[i.Row].Name
			case 1:
				content = fmt.Sprintf("%f", coins[i.Row].Price)
			}
			label.SetText(content)
			column[i.Col] = fyne.Max(column[i.Col], label.MinSize().Width)
		},
	)

	return table
}
