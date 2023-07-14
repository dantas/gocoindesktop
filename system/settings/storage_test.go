package settings

import (
	"os"
	"path"
	"testing"
	"time"

	"github.com/dantas/gocoindesktop/domain"
	"github.com/stretchr/testify/assert"
)

func TestSettingsFileStorageIsCorrectlySavingAndLoading(t *testing.T) {
	// Arrange
	firstStorage := newStorage(t, true)

	settings := domain.Settings{
		Interval:         3 * time.Hour,
		ShowWindowOnOpen: true,
	}

	// Act
	if e := firstStorage.Save(settings); e != nil {
		t.Fatal(e)
	}
	secondStorage := newStorage(t, false)
	loadedSettings, err := secondStorage.Load()

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, settings, loadedSettings)
}

func newStorage(t *testing.T, delete bool) domain.SettingsStorage {
	location := path.Join(os.TempDir(), "settings.json")

	if delete {
		if e := os.Remove(location); e != nil && !os.IsNotExist(e) {
			t.Errorf("Error removing file from temporary storage, test setup error: %v", e)
		}
	}

	return NewSettingsStorage(location)
}
