package cmd

import (
	"fmt"
	"time"

	"github.com/dantas/gocoindesktop/app/settings"
)

func (cl *commandLineController) handleSet(command []string) bool {
	if command[0] != "set" {
		return false
	}

	if command[1] != "interval" {
		return false
	}

	var duration time.Duration
	var err error

	if duration, err = time.ParseDuration(command[2]); err != nil {
		fmt.Println("Error parsing duration", err)
		return false
	}

	settings := settings.Settings{
		Interval: duration,
	}

	if err = cl.application.SetSettings(settings); err != nil {
		fmt.Println("Error saving settings to storage")
	}

	fmt.Println("Interval configured to", duration)

	return true
}

// func handleIntervalSet()
