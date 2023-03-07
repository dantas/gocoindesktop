package cmd

import (
	"fmt"

	"github.com/dantas/gocoindesktop/ui/format"
)

func (cl *commandLineController) handleCoins(command []string) bool {
	if command[0] != "coins" {
		return false
	}

	fmt.Println("Coins:")

	for _, c := range cl.coins {
		fmt.Printf("- %s : %s\n", c.Name, format.FormatPrice(c.Price))
	}

	return true
}