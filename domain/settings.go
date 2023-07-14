package domain

import (
	"errors"
	"time"
)

var (
	ErrSaveSettings = errors.New("error saving settings to disk")
	ErrLoadSettings = errors.New("error loading settings from disk")
)

type Settings struct {
	Interval         time.Duration
	ShowWindowOnOpen bool
}

type SettingsStorage interface {
	Save(Settings) error
	Load() (Settings, error) // If error != nil, returns default settings
}

func NewDefaultSettings() Settings {
	return Settings{
		Interval:         5 * time.Minute,
		ShowWindowOnOpen: true,
	}
}
