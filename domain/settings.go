package domain

import (
	"time"
)

type Settings struct {
	Interval time.Duration
}

var DefaultSettings = Settings{
	Interval: 5 * time.Minute,
}

type SettingsStorage interface {
	Save(Settings) error
	Load() (Settings, error) // In case of error must return (Default.Settings, error)
}
