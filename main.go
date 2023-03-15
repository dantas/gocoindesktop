package main

import (
	"github.com/dantas/gocoindesktop/app/settings"
	"github.com/dantas/gocoindesktop/domain"
	"github.com/dantas/gocoindesktop/domain/alarm"
	"github.com/dantas/gocoindesktop/infrastructure"
	"github.com/dantas/gocoindesktop/ui"
)

func main() {
	application := newApplicationCompositionRoot()
	ui.Run(application)
}

func newApplicationCompositionRoot() *domain.Application {
	settingsStorage := settings.NewJsonFileStorage("settings.json")
	repository := settings.NewRepository(settingsStorage)
	alarmStorage := alarm.NewAlarmStorage("alarm.json")
	alarmManager := alarm.NewAlarmManager(&alarmStorage)
	return domain.NewApplication(domain.NewPeriodicTimer(), repository, infrastructure.CoinMarketCapSource, alarmManager)
}
