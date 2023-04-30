package alarm

import (
	"github.com/dantas/gocoindesktop/domain"
)

type MockStorage struct {
	ToLoadAlarms []domain.Alarm
	SavedAlarms  []domain.Alarm
}

func (mock *MockStorage) Save(alarms []domain.Alarm) error {
	mock.SavedAlarms = alarms
	return nil
}

func (mock *MockStorage) Load() ([]domain.Alarm, error) {
	return mock.ToLoadAlarms, nil
}
