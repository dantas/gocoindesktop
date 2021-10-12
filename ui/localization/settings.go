package localization

var Settings struct {
	UpdateInterval        string
	UpdateIntervalOptions UpdateIntervalOptions
	ShowWindowOnOpen      ShowWindowOnOpen
	SubmitButton          string
}

type ShowWindowOnOpen struct {
	FormLabel   string
	OptionLabel string
}

type UpdateIntervalOptions struct {
	OneMin  string
	TwoMin  string
	FiveMin string
	TenMin  string
	OneHour string
}

func init() {
	Settings.UpdateInterval = "Update interval"
	Settings.ShowWindowOnOpen = ShowWindowOnOpen{
		FormLabel:   "Show window on opening app",
		OptionLabel: "Yes",
	}
	Settings.SubmitButton = "Update"
	Settings.UpdateIntervalOptions = UpdateIntervalOptions{
		"1 min", "2 min", "5 min", "10 min", "1 hour",
	}
}
