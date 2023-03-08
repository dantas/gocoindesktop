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
	settingsManager := domain.NewSettingsManager(settingsStorage)
	return domain.NewApplication(domain.NewPeriodicTimer(), settingsManager, infrastructure.CoinMarketCapSource)
}
