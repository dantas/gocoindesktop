package settings

import (
	"os"
	"path"
	"testing"
	"time"

	"github.com/dantas/gocoindesktop/domain"
)

func TestSettingsFileStorageIsCorrectlySavingAndLoading(t *testing.T) {
	storage := newStorage(t)

	settings := domain.Settings{
		Interval:         3 * time.Hour,
		ShowWindowOnOpen: true,
	}

	if e := storage.Save(settings); e != nil {
		t.Fatal(e)
	}

	if loadedSettings, e := storage.Load(); e != nil {
		t.Error("Error reading file from storage")
	} else if settings != loadedSettings {
		t.Error("Loaded settings is different from what is expected")
	}
}

func TestSettingsFileStorageReturnsDefaultSettingsOnError(t *testing.T) {
	storage := newStorage(t)

	settings, err := storage.Load()

	if err == nil {
		t.Error("Expected error, but found nothing")
	}

	if settings != domain.NewDefaultSettings() {
		t.Errorf("Expected default settings, but found %v", settings)
	}
}

func newStorage(t *testing.T) domain.SettingsStorage {
	location := path.Join(os.TempDir(), "settings.json")

	if e := os.Remove(location); e != nil && !os.IsNotExist(e) {
		t.Errorf("Error removing file from temporary storage, test setup error: %v", e)
	}

	return NewSettingsStorage(location)
}
