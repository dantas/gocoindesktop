package infrastructure

import (
	"os"
	"path"
	"testing"
	"time"

	"github.com/dantas/gocoindesktop/domain"
)

func TestFileSettingsStorageIsCorrectlySavingAndLoading(t *testing.T) {
	location := path.Join(os.TempDir(), "preferences.json")

	if e := os.Remove(location); e != nil {
		t.Errorf("Error removing file from temporary storage, test setup error: %v", e)
	}

	storage := NewJsonFileSettingsStorage(location)

	preference := domain.Settings{
		Interval:         3 * time.Hour,
		ShowWindowOnOpen: true,
	}

	if e := storage.Save(preference); e != nil {
		t.Fatal(e)
	}

	if loadedPreference, e := storage.Load(); e != nil {
		t.Error("Error reading file from storage")
	} else if preference != loadedPreference {
		t.Error("Loaded preference is different from what is expected")
	}
}
