package domain

import (
	"errors"
	"time"
)

var (
	ErrSaveSettings         = errors.New("error saving settings to disk")
	ErrLoadSettings         = errors.New("error loading settings from disk")
	ErrLoadSettingsNotExist = errors.New("error persisted settings do not exist")
)

type Settings struct {
	Interval         time.Duration
	ShowWindowOnOpen bool
}

type SettingsStorage interface {
	Save(Settings) error
	Load() (Settings, error) // If error != nil, returns default settings
}

func newDefaultSettings() Settings {
	return Settings{
		Interval:         5 * time.Minute,
		ShowWindowOnOpen: true,
	}
}
