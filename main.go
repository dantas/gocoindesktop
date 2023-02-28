package main

import (
	"github.com/dantas/gocoindesktop/domain"
	"github.com/dantas/gocoindesktop/infrastructure"
	"github.com/dantas/gocoindesktop/ui/gui"
)

func main() {
	application := newApplicationCompositionRoot()
	gui.RunGui(application)
}

func newApplicationCompositionRoot() *domain.Application {
	settingsStorage := infrastructure.NewJsonFileSettingsStorage("settings.json")
	coinTicker := domain.NewCoinTicker(infrastructure.CoinMarketCapSource)
	return domain.NewApplication(coinTicker, settingsStorage)
}
