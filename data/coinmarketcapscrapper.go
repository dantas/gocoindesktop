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

	go executeScrap(ctx, resultChannel)

	return resultChannel
}

func executeScrap(ctx context.Context, resultChannel chan domain.ScrapResult) {
	isActive := true

	close := func() {
		isActive = false
		close(resultChannel)
	}

	closeIfNecessary := func() bool {
		if !isActive {
			return true
		}

		select {
		case <-ctx.Done():
			close()
			return true
		default:
			return false
		}
	}

	coins := make([]domain.Coin, 0)
	errors := make([]error, 0)

	collector := colly.NewCollector()

	collector.OnResponse(func(r *colly.Response) {
		// sleep one second to ensure javascript has time to update the listing
		time.Sleep(1 * time.Second)
	})

	collector.OnHTML(`tr[class="cmc-table-row"]`, func(e *colly.HTMLElement) {
		if closeIfNecessary() {
			return
		}

		scrappedName := e.ChildAttr("a", "title")
		scrappedPrice := e.ChildText(`td:nth-child(5) a`)

		sanitizedPrice := strings.Replace(scrappedPrice[1:], ",", "", -1)
		priceFloat, err := strconv.ParseFloat(sanitizedPrice, 64)

		if err != nil {
			errors = append(errors, err)
		} else {
			coins = append(
				coins,
				domain.Coin{
					Name:  scrappedName,
					Price: priceFloat,
				},
			)
		}
	})

	sendResultsAndClose := func() {
		result := domain.ScrapResult{
			Coins:  coins,
			Errors: errors,
		}

		select {
		case resultChannel <- result:
			break
		case <-ctx.Done():
			break
		}

		isActive = false

		close()
	}

	collector.OnScraped(func(r *colly.Response) {
		if closeIfNecessary() {
			return
		}

		sendResultsAndClose()
	})

	collector.OnError(func(r *colly.Response, e error) {
		if closeIfNecessary() {
			return
		}

		errors = append(errors, e)

		sendResultsAndClose()
	})

	collector.Visit("https://coinmarketcap.com/all/views/all/")

	collector.Wait()
}
