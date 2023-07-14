package alarm

import (
	"os"
	"path"
	"testing"

	"github.com/dantas/gocoindesktop/domain"
	"github.com/stretchr/testify/assert"
)

func TestSettingsFileStorageIsCorrectlySavingAndLoading(t *testing.T) {
	// Arrange
	firstStorage := newStorage(t, true)

	alarms := []domain.Alarm{
		{
			Name:       "First Coin",
			LowerBound: 1500,
			UpperBound: 3000,
			IsEnabled:  true,
		},
		{
			Name:       "Second Coin",
			LowerBound: 2500,
			UpperBound: 3500,
			IsEnabled:  false,
		},
	}

	// Act
	if e := firstStorage.Save(alarms); e != nil {
		t.Fatal(e)
	}

	secondStorage := newStorage(t, false)
	loadedAlarms, err := secondStorage.Load()

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, alarms, loadedAlarms)
}

func TestIfFileDoesntExistReturnExpectedError(t *testing.T) {
	// Arrange
	storage := newStorage(t, true)

	// Act
	content, err := storage.Load()

	// Assert
	assert.Nil(t, content)
	assert.Equal(t, domain.ErrLoadAlarmNotExist, err)
}

func newStorage(t *testing.T, delete bool) domain.AlarmStorage {
	location := path.Join(os.TempDir(), "alarms.json")

	if delete {
		if e := os.Remove(location); e != nil && !os.IsNotExist(e) {
			t.Errorf("Error removing file from temporary storage, test setup error: %v", e)
		}
	}

	return NewAlarmStorage(location)
}
