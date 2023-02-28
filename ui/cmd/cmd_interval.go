package cmd

import "fmt"

func (cl *commandLineController) handleInterval(command []string) bool {
	if command[0] != "interval" {
		return false
	}

	intervalStr := cl.application.Settings().Interval

	fmt.Println("Refresh interval is", intervalStr)

	return true
}
