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
		ShowWindowOnOpen: true,
	}
}

// Find better name
type settingsManager struct {
	settings Settings
	storage  SettingsStorage
}

func NewSettingsManager(storage SettingsStorage) *settingsManager {
	return &settingsManager{
		storage: storage,
	}
}

func (m *settingsManager) Load() error {
	var err error

	if settings, err := m.storage.Load(); err != nil {
		m.settings = newDefaultSettings()
	} else {
		m.settings = settings
	}

	return err
}

func (m *settingsManager) Save(settings Settings) error {
	if e := m.storage.Save(settings); e != nil {
		return e
	}

	m.settings = settings

	return nil
}
