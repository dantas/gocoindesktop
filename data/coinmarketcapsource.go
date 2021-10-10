package data

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/dantas/gocoindesktop/domain"
	"github.com/gocolly/colly/v2"
)

func CoinMarketCapSource(ctx context.Context) <-chan domain.CoinSourceResult {
	chanResults := make(chan domain.CoinSourceResult)

	go fetchCoins(ctx, chanResults)

	return chanResults
}

func fetchCoins(ctx context.Context, chanResults chan domain.CoinSourceResult) {
	isActive := true

	sendResult := func(result domain.CoinSourceResult) {
		select {
		case <-ctx.Done():
			isActive = false
			close(chanResults)
		case chanResults <- result:
		}
	}

	collector := colly.NewCollector()

	collector.OnResponse(func(r *colly.Response) {
		// sleep one second to ensure javascript has time to update the listing
		time.Sleep(1 * time.Second)
	})

	collector.OnHTML(`tr[class="cmc-table-row"]`, func(e *colly.HTMLElement) {
		if !isActive {
			return
		}

		scrappedName := e.ChildAttr("a", "title")
		scrappedPrice := e.ChildText(`td:nth-child(5) a`)

		sanitizedPrice := strings.Replace(scrappedPrice[1:], ",", "", -1)
		priceFloat, err := strconv.ParseFloat(sanitizedPrice, 64)

		result := domain.CoinSourceResult{}

		if err != nil {
			result.Error = err
		} else {
			result.Coin = domain.Coin{
				Name:  scrappedName,
				Price: priceFloat,
			}
		}

		sendResult(result)
	})

	collector.OnScraped(func(r *colly.Response) {
		if isActive {
			close(chanResults)
		}
	})

	collector.OnError(func(r *colly.Response, err error) {
		if !isActive {
			return
		}

		result := domain.CoinSourceResult{
			Error: err,
		}

		sendResult(result)
	})

	collector.Visit("https://coinmarketcap.com/all/views/all/")

	collector.Wait()
}
