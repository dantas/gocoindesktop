package main

import (
	"github.com/dantas/gocoindesktop/domain"
	"github.com/dantas/gocoindesktop/infrastructure"
	"github.com/dantas/gocoindesktop/ui"
)

func main() {
	application := newApplicationCompositionRoot()
	ui.Run(application)
}

func newApplicationCompositionRoot() *domain.Application {
	settingsStorage := infrastructure.NewJsonFileSettingsStorage("settings.json")
	coinTicker := domain.NewCoinTicker(infrastructure.CoinMarketCapSource)
	return domain.NewApplication(coinTicker, settingsStorage)
}
