package domain

import (
	"context"
)

type Scrapper func(ctx context.Context) <-chan ScrapResult

type ScrapResult struct {
	Coins  []Coin
	Errors []error
}
