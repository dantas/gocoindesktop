package data

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/dantas/gocoindesktop/domain"
	"github.com/gocolly/colly/v2"
)

func CoinMarketCapScrapper(ctx context.Context) <-chan domain.ScrapResult {
	resultChannel := make(chan domain.ScrapResult)

	go func() {
		collector := colly.NewCollector()

		collector.OnResponse(func(r *colly.Response) {
			// sleep one second to ensure javascript has time to update the listing
			time.Sleep(1 * time.Second)
		})

		isActive := true

		collector.OnHTML(`tr[class="cmc-table-row"]`, func(e *colly.HTMLElement) {
			if !isActive {
				return
			}

			var result domain.ScrapResult

			scrappedName := e.ChildAttr("a", "title")
			scrappedPrice := e.ChildText(`td:nth-child(5) a`)

			sanitizedPrice := strings.Replace(scrappedPrice[1:], ",", "", -1)
			priceFloat, error := strconv.ParseFloat(sanitizedPrice, 64)
			if error != nil {
				result = domain.ScrapResult{Error: error}
			} else {
				result = domain.ScrapResult{
					Coin: domain.Coin{
						Name:  scrappedName,
						Price: priceFloat,
					},
				}
			}

			select {
			case resultChannel <- result:
				break
			case <-ctx.Done():
				close(resultChannel)
				isActive = false
				break
			}
		})

		collector.OnScraped(func(r *colly.Response) {
			if isActive {
				close(resultChannel)
			}
		})

		collector.OnError(func(r *colly.Response, e error) {
			if !isActive {
				return
			}

			isActive = false

			resultChannel <- domain.ScrapResult{
				Error: e,
			}

			close(resultChannel)
		})

		collector.Visit("https://coinmarketcap.com/all/views/all/")

		collector.Wait()
	}()

	return resultChannel
}
