package main

import (
	"github.com/dantas/gocoindesktop/domain"
	"github.com/dantas/gocoindesktop/system/alarm"
	"github.com/dantas/gocoindesktop/system/coinsource"
	"github.com/dantas/gocoindesktop/system/settings"
	"github.com/dantas/gocoindesktop/system/timer"
	"github.com/dantas/gocoindesktop/ui/fynegui"
)

func main() {
	application := newApplicationCompositionRoot()
	fynegui.Run(application)
}

func newApplicationCompositionRoot() *domain.Application {
	settingsStorage := settings.NewSettingsStorage("settings.json")
	alarmStorage := alarm.NewAlarmStorage("alarms.json")
	alarmManager := domain.NewAlarmManager(alarmStorage)
	return domain.NewApplication(timer.NewPeriodicTimer(), settingsStorage, coinsource.CoinMarketCapSource, alarmManager)
}
