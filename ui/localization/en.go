//go:build !pt

package localization

import (
	"fmt"

	"github.com/dantas/gocoindesktop/app/alarm"
)

// Systray
const SystrayCoins = "Show coins"
const SystraySettings = "Show settings"
const SystrayQuit = "Quit"

// Window
const AppTitle = "Go Coin Desktop"

// Window Menu
const TabCoins = "Coins"
const TabSettings = "Settings"

// Coin Tab
const ColumnCoin = "Coin"
const ColumnPrice = "Price"
const ColumnAlarm = "Enable alarm"
const ColumnLowerBound = "Lower bound"
const ColumnUpperBound = "Upper bound"

// Alarm notification
func AlarmTitle(alarm alarm.TriggeredAlarm) string {
	return fmt.Sprintf("Alert for %s", alarm.Coin.Name)
}

func AlarmEnterRangeMessage(alarm alarm.TriggeredAlarm) string {
	return fmt.Sprintf("Coin %s in range, price: %s", alarm.Coin.Name, FormatPrice(alarm.Coin.Price))
}

func AlarmLeaveRangeMessage(alarm alarm.TriggeredAlarm) string {
	return fmt.Sprintf("Coin %s out of range, price: %s", alarm.Coin.Name, FormatPrice(alarm.Coin.Price))
}

// Settings tab
const SettingsUpdateInterval = "Update interval"
const SettingsShowWindowOnOpen = "Show window on opening app"
const SettingsShowWindowOnOpenOption = "Yes"
const SettingsSubmitButton = "Update"
const Settings1Min = "1 min"
const Settings2Min = "2 min"
const Settings5Min = "5 min"
const Settings10Min = "10 min"
const Settings1Hour = "1 hour"
