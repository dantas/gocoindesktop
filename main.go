package main

import (
	"github.com/dantas/gocoindesktop/app"
	"github.com/dantas/gocoindesktop/app/alarm"
	"github.com/dantas/gocoindesktop/app/coin"
	"github.com/dantas/gocoindesktop/app/settings"
	"github.com/dantas/gocoindesktop/app/timer"
	"github.com/dantas/gocoindesktop/ui/gui"
)

func main() {
	application := newApplicationCompositionRoot()
	gui.RunFyne(application)
}

func newApplicationCompositionRoot() *app.Application {
	settingsStorage := settings.NewJsonFileStorage("settings.json")
	repository := settings.NewRepository(settingsStorage)
	alarmStorage := alarm.NewAlarmStorage("alarm.json")
	alarmManager := alarm.NewAlarmManager(&alarmStorage)
	return app.NewApplication(timer.NewPeriodicTimer(), repository, coin.CoinMarketCapSource, alarmManager)
}
