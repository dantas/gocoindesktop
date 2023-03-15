package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dantas/gocoindesktop/app"
	"github.com/dantas/gocoindesktop/app/coin"
)

func RunCmd(application *app.Application) {
	// return from command structure to print it here
	fmt.Println("Welcome to Go Coin Desktop")
	fmt.Println("Available commands:")
	fmt.Println("- coins: Show coins and prices")
	fmt.Println("- interval: Show update interval")
	fmt.Println("- interval 1 min: Set update interval")
	fmt.Println("- quit: Quit application")

	// TODO: Talvez n seja necesario isto
	controller := commandLineController{
		application: application,
	}

	go func() {
		for coins := range application.Coins() {
			controller.coins = coins
		}
	}()

	for command := range readCommands("quit") {
		if controller.handleCoins(command) {
			continue
		}

		if controller.handleInterval(command) {
			continue
		}

		if controller.handleSet(command) {
			continue
		}

		fmt.Println("Invalid command")
	}

	application.Destroy()
}

type commandLineController struct {
	coins       []coin.Coin
	application *app.Application
}

// Quit command needs be handled here, because is nigh impossible to interrupt a scanner
func readCommands(quitCommand string) <-chan []string {
	channel := make(chan []string)

	lineScanner := bufio.NewScanner(bufio.NewReader(os.Stdin))

	go func() {
		defer close(channel)

		for lineScanner.Scan() {
			words := strings.Fields(lineScanner.Text())

			switch {
			case len(words) == 0:
				continue
			case words[0] == quitCommand:
				return
			default:
				channel <- words
			}
		}
	}()

	return channel
}
