package settings

import (
	"github.com/dantas/gocoindesktop/domain"
)

type MockStorage struct {
	SavedSettings  domain.Settings
	ToLoadSettings domain.Settings
}

func (mock *MockStorage) Save(settings domain.Settings) error {
	mock.SavedSettings = settings
	return nil
}

func (mock *MockStorage) Load() (domain.Settings, error) {
	return mock.ToLoadSettings, nil
}
