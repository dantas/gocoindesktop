package data

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dantas/gocoindesktop/domain"
	"github.com/gocolly/colly/v2"
)

/*
	TODO:
	- Return errors
	- Filter by name
*/
func ScrapCoinMarketCap(done <-chan interface{}) <-chan domain.Coin {
	coinChannel := make(chan domain.Coin)

	go func() {
		collector := colly.NewCollector()

		collector.OnResponse(func(r *colly.Response) {
			// sleep one second to ensure javascript has time to update the listing
			var d, _ = time.ParseDuration("1s")
			time.Sleep(d)
		})

		isActive := true

		collector.OnHTML(`tr[class="cmc-table-row"]`, func(e *colly.HTMLElement) {
			if !isActive {
				return
			}

			fmt.Println("Scrapping")

			scrappedName := e.ChildAttr("a", "title")
			scrappedPrice := e.ChildText(`td:nth-child(5) a`)

			sanitizedPrice := strings.Replace(scrappedPrice[1:], ",", "", -1)
			priceFloat, error := strconv.ParseFloat(sanitizedPrice, 64)
			if error != nil {
				panic(error)
			}

			coin := domain.Coin{
				Name:  scrappedName,
				Price: priceFloat,
			}

			select {
			case coinChannel <- coin:
				break
			case <-done:
				close(coinChannel)
				isActive = false
				break
			}
		})

		collector.OnScraped(func(r *colly.Response) {
			if isActive {
				close(coinChannel)
			}
		})

		collector.Visit("https://coinmarketcap.com/all/views/all/")

		collector.Wait()
	}()

	return coinChannel
}
