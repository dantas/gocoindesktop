package localization

var Settings struct {
	UpdateInterval        string
	SubmitButton          string
	UpdateIntervalOptions UpdateIntervalOptions
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
	Settings.SubmitButton = "Update"
	Settings.UpdateIntervalOptions = UpdateIntervalOptions{
		"1 min", "2 min", "5 min", "10 min", "1 hour",
	}
}
