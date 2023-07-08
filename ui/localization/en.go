//go:build !pt

package localization

import (
	"fmt"

	"fyne.io/fyne/v2"
	"github.com/dantas/gocoindesktop/domain"
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

// Sizes
func WindowSize() fyne.Size {
	return fyne.NewSize(600, 300)
}

const ColumnWidthCoin = 150
const ColumnWidthPrice = 100
const ColumnWidthAlarm = 100
const ColumnWidthLowerBound = 100
const ColumnWidthUpperBound = 100

// Alarm notification
func AlarmTitle(alarm domain.TriggeredAlarm) string {
	return fmt.Sprintf("%s Alert", alarm.Coin.Name)
}

func AlarmEnterRangeMessage(alarm domain.TriggeredAlarm) string {
	return fmt.Sprintf("Coin in range, price: %s", FormatPrice(alarm.Coin.Price))
}

func AlarmLeaveRangeMessage(alarm domain.TriggeredAlarm) string {
	return fmt.Sprintf("Coin out of range, price: %s", FormatPrice(alarm.Coin.Price))
}

// Settings tab
const SettingsUpdateInterval = "Update interval"
const SettingsShowWindowOnOpen = "Show window when opening app"
const SettingsShowWindowOnOpenOption = "Yes"
const Settings1Min = "1 min"
const Settings2Min = "2 min"
const Settings5Min = "5 min"
const Settings10Min = "10 min"
const Settings1Hour = "1 hour"
