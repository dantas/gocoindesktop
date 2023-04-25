package alarm

import (
	"os"
	"path"
	"testing"

	"github.com/dantas/gocoindesktop/domain"
	"github.com/dantas/gocoindesktop/utils"
)

func TestSettingsFileStorageIsCorrectlySavingAndLoading(t *testing.T) {
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

	if e := firstStorage.Save(alarms); e != nil {
		t.Fatal(e)
	}

	secondStorage := newStorage(t, false)

	if loadedAlarms, e := secondStorage.Load(); e != nil {
		t.Error("Error reading file from storage")
	} else if !utils.Equals(alarms, loadedAlarms) {
		t.Error("Loaded settings is different from what is expected")
	}
}

func TestIfFileDoesntExistExpectError(t *testing.T) {
	storage := newStorage(t, true)

	content, err := storage.Load()

	if content != nil {
		t.Error("File doesn't exist but storage decided to return something")
	}

	if err == nil {
		t.Error("File doesn't exist but storage returned no error")
	}
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
