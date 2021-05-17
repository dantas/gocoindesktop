package main

import (
	"fmt"

	"github.com/dantas/gocoindesktop/internal/scrapper"
)

func main() {
	var result = scrapper.Scrap()

	for _, c := range result {
		fmt.Printf("%s : %.2f\n", c.Name, c.Value)
	}
}
