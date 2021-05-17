package scrapper

import (
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

func Scrap() []Coin {
	var coins = make([]Coin, 0)

	collector := colly.NewCollector()

	collector.OnResponse(func(r *colly.Response) {
		// sleep one second to ensure javascript has time to update the listing
		var d, _ = time.ParseDuration("1s")
		time.Sleep(d)
	})

	// extract names
	collector.OnHTML(`p[class="sc-1eb5slv-0 iJjGCS"]`, func(h *colly.HTMLElement) {
		coins = append(coins, Coin{
			Name: h.Text,
		})
	})

	// extract values
	var index uint
	collector.OnHTML(`div[class="price___3rj7O "]`, func(h *colly.HTMLElement) {
		var htmlText = h.ChildText("a")
		var toParseText = strings.Replace(htmlText[1:], ",", "", -1)

		var float, error = strconv.ParseFloat(toParseText, 64)
		if error != nil {
			panic(error)
		}

		coins[index].Value = float

		index++
	})

	collector.Visit("https://coinmarketcap.com/en/")

	collector.Wait()

	return coins
}
