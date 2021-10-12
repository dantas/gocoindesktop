package domain

import (
	"time"
)

type Settings struct {
	Interval         time.Duration
	ShowWindowOnOpen bool
}

var DefaultSettings = Settings{
	Interval:         5 * time.Minute,
	ShowWindowOnOpen: false,
}

type SettingsStorage interface {
	Save(Settings) error
	Load() (Settings, error) // In case of error must return (Default.Settings, error)
}
