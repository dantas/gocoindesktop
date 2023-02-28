package cmd

import "fmt"

func (cl *commandLineController) handleCoins(command []string) bool {
	if command[0] != "coins" {
		return false
	}

	fmt.Println("Coins:")

	for _, c := range cl.coins {
		fmt.Printf("- %s : %f\n", c.Name, c.Price)
	}

	return true
}
