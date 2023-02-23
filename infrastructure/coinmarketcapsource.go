package infrastructure

import (
	"strconv"
	"strings"
	"time"

	"github.com/dantas/gocoindesktop/domain"
	"github.com/gocolly/colly/v2"
)

func CoinMarketCapSource() ([]domain.Coin, error) {
	var resultCoins = make([]domain.Coin, 0)
	var resultError error

	collector := colly.NewCollector()

	collector.OnResponse(func(r *colly.Response) {
		// sleep one second to ensure javascript has time to update the listing
		time.Sleep(1 * time.Second)
	})

	collector.OnHTML(`tr[class="cmc-table-row"]`, func(e *colly.HTMLElement) {
		scrappedName := e.ChildAttr("a", "title")
		scrappedPrice := e.ChildText(`td:nth-child(5) a`)

		sanitizedPrice := strings.Replace(scrappedPrice[1:], ",", "", -1)
		priceFloat, err := strconv.ParseFloat(sanitizedPrice, 64)

		if err != nil {
			resultError = err
		} else {
			resultCoins = append(resultCoins, domain.Coin{
				Name:  scrappedName,
				Price: priceFloat,
			})
		}
	})

	collector.OnError(func(r *colly.Response, err error) {
		resultError = err
	})

	collector.Visit("https://coinmarketcap.com/all/views/all/")

	collector.Wait()

	return resultCoins, resultError
}
