//go:build pt

package localization

import (
	"fmt"

	"github.com/dantas/gocoindesktop/domain/alarm"
)

// Systray
const SystrayCoins = "Mostrar moedas"
const SystraySettings = "Mostrar configurações"
const SystrayQuit = "Sair"

// Window
const AppTitle = "Go Coin Desktop"

// Window Menu
const TabCoins = "Moedas"
const TabSettings = "Configurações"

// Coin Tab
const ColumnCoin = "Moedas"
const ColumnPrice = "Preço"
const ColumnAlarm = "Habilitar alarme"
const ColumnLowerBound = "Limite inferior"
const ColumnUpperBound = "Limite superior"

// Alarm notification
func AlarmTitle(alarm alarm.TriggeredAlarm) string {
	return fmt.Sprintf("Alerta para %s", alarm.Coin.Name)
}

func AlarmEnterRangeMessage(alarm alarm.TriggeredAlarm) string {
	return fmt.Sprintf("Moeda %s dentro do intervalo, preço: %s", alarm.Coin.Name, FormatPrice(alarm.Coin.Price))
}

func AlarmLeaveRangeMessage(alarm alarm.TriggeredAlarm) string {
	return fmt.Sprintf("Moeda %s fora do intervalo, preço: %s", alarm.Coin.Name, FormatPrice(alarm.Coin.Price))
}

// Settings tab
const SettingsUpdateInterval = "Atualizar intervalo"
const SettingsShowWindowOnOpen = "Mostrar janela na abertura"
const SettingsShowWindowOnOpenOption = "Sim"
const SettingsSubmitButton = "Atualizar"
const Settings1Min = "1 min"
const Settings2Min = "2 min"
const Settings5Min = "5 min"
const Settings10Min = "10 min"
const Settings1Hour = "1 hour"
