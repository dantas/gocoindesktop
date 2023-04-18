package main

import (
	"github.com/dantas/gocoindesktop/app"
	"github.com/dantas/gocoindesktop/app/alarm"
	"github.com/dantas/gocoindesktop/app/coin"
	"github.com/dantas/gocoindesktop/app/settings"
	"github.com/dantas/gocoindesktop/app/timer"
	"github.com/dantas/gocoindesktop/ui/fynegui"
)

func main() {
	application := newApplicationCompositionRoot()
	fynegui.Run(application)
}

func newApplicationCompositionRoot() *app.Application {
	settingsStorage := settings.NewSettingsStorage("settings.json")
	alarmStorage := alarm.NewAlarmStorage("alarms.json")
	alarmManager := alarm.NewAlarmManager(alarmStorage)
	return app.NewApplication(timer.NewPeriodicTimer(), settingsStorage, coin.CoinMarketCapSource, alarmManager)
}
