package localization

var Systray struct {
	Coins    string
	Settings string
	Quit     string
}

func init() {
	Systray.Coins = "Show coins"
	Systray.Settings = "Show settings"
	Systray.Quit = "Quit"
}
