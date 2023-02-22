package domain

import (
	"time"
)

type Settings struct {
	Interval         time.Duration
	ShowWindowOnOpen bool
}

type SettingsStorage interface {
	Save(Settings) error
	Load() (Settings, error)
}

func newDefaultSettings() Settings {
	return Settings{
		Interval:         5 * time.Minute,
		ShowWindowOnOpen: false,
	}
}
